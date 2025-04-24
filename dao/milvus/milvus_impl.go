package milvus

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

type milvusDao struct {
	milvusCli client.Client
}

func NewMilvusDao(cli client.Client) Dao {
	return &milvusDao{milvusCli: cli}
}

func (m *milvusDao) SaveChunks(ctx context.Context, chunks []Chunk) error {
	if len(chunks) == 0 {
		return fmt.Errorf("num_rows should be greater than 0: invalid parameter[expected=invalid num_rows][actual=0")
	}

	// 准备有效的数据
	preparedData, err := m.prepareChunkData(chunks)
	if err != nil {
		return err
	}

	// 创建数据列
	columns := m.createDataColumns(preparedData)

	// 插入数据
	return m.insertDataWithRetry(ctx, preparedData.collectionName, columns, 3)
}

// prepareChunkData 验证和准备文本块数据
// 该方法会过滤掉无效的文本块（空内容或空向量），并确保文档名不超过最大长度限制
// 返回处理后的数据结构和可能的错误
func (m *milvusDao) prepareChunkData(chunks []Chunk) (*chunkData, error) {
	milvusConfig := conf.GetCfg().Milvus
	data := &chunkData{
		collectionName: CollectionNameTextChunks,
		vectorDim:      milvusConfig.VectorDimension,
	}

	// 遍历验证并准备数据
	for i, chunk := range chunks {
		// 验证chunk数据
		if len(chunk.Content) == 0 {
			fmt.Printf("prepareChunkData warn: No.%d vector is null, it will be passed\n", i)
			continue
		}
		if len(chunk.Embeddings) == 0 {
			fmt.Printf("prepareChunkData warn: No.%d vector is null, it will be passed\n", i)
			continue
		}

		// 确保文档名不超过限制长度
		docName := chunk.DocumentName
		if len(docName) > 250 {
			docName = docName[:250]
		}

		// 添加有效数据
		data.ids = append(data.ids, chunk.ID)
		data.contents = append(data.contents, chunk.Content)
		data.documentIDs = append(data.documentIDs, chunk.DocumentID)
		data.documentNames = append(data.documentNames, docName)
		data.kbIDs = append(data.kbIDs, chunk.KBID)
		data.chunkIndices = append(data.chunkIndices, int32(chunk.Index))
		data.vectors = append(data.vectors, chunk.Embeddings)
	}

	// 确保有有效数据要插入
	if len(data.ids) == 0 {
		return nil, fmt.Errorf("过滤无效数据后，没有有效的文本块可以插入")
	}
	return data, nil
}

// createDataColumns 创建Milvus数据列
// 将预处理的数据转换为Milvus插入操作所需的列格式
// 返回包含所有数据列的切片
func (m *milvusDao) createDataColumns(data *chunkData) []entity.Column {
	idColumn := entity.NewColumnVarChar(FieldNameID, data.ids)
	contentColumn := entity.NewColumnVarChar(FieldNameContent, data.contents)
	documentIDColumn := entity.NewColumnVarChar(FieldNameDocumentID, data.documentIDs)
	documentNameColumn := entity.NewColumnVarChar(FieldNameDocumentName, data.documentNames)
	kbIDColumn := entity.NewColumnVarChar(FieldNameKBID, data.kbIDs)
	indexColumn := entity.NewColumnInt32(FieldNameChunkIndex, data.chunkIndices)
	vectorColumn := entity.NewColumnFloatVector(FieldNameVector, data.vectorDim, data.vectors)

	return []entity.Column{
		idColumn,
		contentColumn,
		documentIDColumn,
		documentNameColumn,
		kbIDColumn,
		indexColumn,
		vectorColumn,
	}
}

func (m *milvusDao) insertDataWithRetry(ctx context.Context, collectionName string,
	columns []entity.Column, maxRetries int) error {
	var result *multierror.Error
	baseDelay := 100 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		fmt.Printf("insertDataWithRetry debug, try to insert data: (%d/%d)...\n", i+1, maxRetries)
		_, err := m.milvusCli.Insert(ctx, collectionName, "", columns...)
		if err == nil {
			fmt.Printf("insertDataWithRetry Success!\n")
			return nil
		}

		// 记录结果
		result = multierror.Append(result, fmt.Errorf("insertDataWithRetry %d/%d failed: %w", i+1, maxRetries, err))

		// 指数退避
		delay := baseDelay * time.Duration(1<<uint(i))
		time.Sleep(delay)
	}
	return result.ErrorOrNil()
}

// DeleteChunks 删除指定文档ID列表对应的所有文本块
// 使用IN操作符构建删除表达式，一次性删除多个文档的所有块
func (m *milvusDao) DeleteChunks(ctx context.Context, docIDs []string) error {
	// 构建删除表达式，使用 IN 操作符
	expr := fmt.Sprintf("%s in [\"%s\"]", FieldNameDocumentID, strings.Join(docIDs, "\",\""))
	// 删除
	if err := m.milvusCli.Delete(ctx, CollectionNameTextChunks, "", expr); err != nil {
		return fmt.Errorf("删除向量数据失败：%w", err)
	}

	return nil
}

// Search 在知识库中搜索相似向量
// 参数:
//   - kbID: 知识库ID，指定搜索范围
//   - vector: 查询向量，通常是问题或查询文本的嵌入表示
//   - topK: 返回的最相似结果数量
//
// 返回:
//   - 按相似度排序的文本块切片
//   - 可能的错误
func (m *milvusDao) Search(ctx context.Context, kbID string, vector []float32, topK int) ([]Chunk, error) {
	// 构建搜索参数
	sp, _ := entity.NewIndexIvfFlatSearchParam(conf.GetCfg().Milvus.Nprobe)
	expr := fmt.Sprintf("%s == \"%s\"", FieldNameKBID, kbID)

	// 执行搜索
	searchResult, err := m.milvusCli.Search(
		ctx,
		CollectionNameTextChunks, // 要搜索的集合名称
		[]string{},               // 要返回的字段名数组，空数组表示返回所有字段
		expr,                     // 过滤表达式，用于筛选符合条件的文档
		SearchFields,             // 要搜索的字段名称列表
		[]entity.Vector{entity.FloatVector(vector)}, // 搜索向量，将查询文本转换的向量用于相似度搜索
		FieldNameVector,                      // 向量字段的名称，指定哪个字段存储的是向量数据
		conf.GetCfg().Milvus.GetMetricType(), // 相似度度量类型（如欧氏距离、余弦相似度等）
		topK,                                 // 返回的最相似结果数量
		sp,                                   // 搜索参数，包含额外的搜索配置选项
	)
	if err != nil {
		return nil, fmt.Errorf("向量检索失败: %w", err)
	}
	return m.parseSearchResults(searchResult)
}

// parseSearchResults 解析搜索结果，将Milvus返回结果转换为模型数据
// 参数:
//   - searchResult: Milvus搜索结果
//
// 返回:
//   - 解析后的文本块切片，按相似度得分降序排序
//   - 可能的错误
func (m *milvusDao) parseSearchResults(searchResult []client.SearchResult) ([]Chunk, error) {
	var chunks []Chunk
	// log.Printf("SearchResult长度：%v\n", len(searchResult))
	// for _, res := range searchResult {
	// 	log.Printf("IDs: %s\n", res.IDs)
	// 	log.Printf("Fields: %s\n", res.Fields)
	// 	log.Printf("Scores: %v\n", res.Scores)
	// }

	for _, result := range searchResult {
		idCol, ok := result.IDs.(*entity.ColumnVarChar)
		if !ok {
			return nil, fmt.Errorf("unexpected type for ID column: %T", result.IDs)
		}

		contentCol, ok := result.Fields.GetColumn(FieldNameContent).(*entity.ColumnVarChar)
		if !ok {
			return nil, fmt.Errorf("unexpected type for content column")
		}

		docIDCol, ok := result.Fields.GetColumn(FieldNameDocumentID).(*entity.ColumnVarChar)
		if !ok {
			return nil, fmt.Errorf("unexpected type for document ID column")
		}

		docNameCol, ok := result.Fields.GetColumn(FieldNameDocumentName).(*entity.ColumnVarChar)
		if !ok {
			return nil, fmt.Errorf("unexpected type for document Name column")
		}

		kbIDCol, ok := result.Fields.GetColumn(FieldNameKBID).(*entity.ColumnVarChar)
		if !ok {
			return nil, fmt.Errorf("unexpected type for KB ID column")
		}

		indexCol, ok := result.Fields.GetColumn(FieldNameChunkIndex).(*entity.ColumnInt32)
		if !ok {
			return nil, fmt.Errorf("unexpected type for index column")
		}

		resultCount := idCol.Len()
		for i := 0; i < resultCount; i++ {
			id := idCol.Data()[i]
			content, err := contentCol.GetAsString(i)
			if err != nil {
				return nil, fmt.Errorf("获取内容失败: %w", err)
			}

			docID, err := docIDCol.GetAsString(i)
			if err != nil {
				return nil, fmt.Errorf("获取文档ID失败: %w", err)
			}

			docName, err := docNameCol.GetAsString(i)
			if err != nil {
				return nil, fmt.Errorf("获取文档名称失败：%w", err)
			}

			kbID, err := kbIDCol.GetAsString(i)
			if err != nil {
				return nil, fmt.Errorf("获取知识库ID失败: %w", err)
			}

			index := indexCol.Data()[i]

			score := result.Scores[i]

			chunks = append(chunks, Chunk{
				ID:           id,
				Content:      content,
				KBID:         kbID,
				DocumentID:   docID,
				DocumentName: docName,
				Index:        int(index),
				Score:        score,
			})
		}
	}

	// 按Score从高到低排序
	sort.Slice(chunks, func(i, j int) bool {
		return chunks[i].Score > chunks[j].Score
	})
	return chunks, nil
}
