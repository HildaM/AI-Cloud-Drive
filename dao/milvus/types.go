package milvus

// 存储到milvus中
type Chunk struct {
	ID           string    `json:"id"`
	Content      string    `json:"content"`       // chunk内容
	KBID         string    `json:"kb_id"`         // 知识库ID（知识库级别的检索）
	DocumentID   string    `json:"document_id"`   // 文档ID
	DocumentName string    `json:"document_name"` // 文档名称
	Index        int       `json:"index"`         // 第几个chunk
	Embeddings   []float32 `json:"embeddings"`    // chunk向量
	Score        float32   `json:"score"`         // 返回分数信息
}

// chunkData 文本块数据结构，用于存储预处理后的向量数据
type chunkData struct {
	collectionName string      // 目标集合名称
	vectorDim      int         // 向量维度
	ids            []string    // 文本块ID列表
	contents       []string    // 文本块内容列表
	documentIDs    []string    // 对应的文档ID列表
	documentNames  []string    // 对应的文档名称列表
	kbIDs          []string    // 对应的知识库ID列表
	chunkIndices   []int32     // 文本块在文档中的索引位置
	vectors        [][]float32 // 文本块的向量表示
}

// 集合相关常量定义
const (
	// CollectionNameTextChunks 文本块集合名称
	CollectionNameTextChunks = "text_chunks"
)

// 字段名称常量定义
const (
	// FieldNameID ID字段名
	FieldNameID = "id"
	// FieldNameContent 内容字段名
	FieldNameContent = "content"
	// FieldNameDocumentID 文档ID字段名
	FieldNameDocumentID = "document_id"
	// FieldNameDocumentName 文档名称字段名
	FieldNameDocumentName = "document_name"
	// FieldNameKBID 知识库ID字段名
	FieldNameKBID = "kb_id"
	// FieldNameChunkIndex 块索引字段名
	FieldNameChunkIndex = "chunk_index"
	// FieldNameVector 向量字段名
	FieldNameVector = "vector"
)

// 查询相关字段
var (
	// SearchFields 搜索结果返回的字段
	SearchFields = []string{
		FieldNameID,
		FieldNameContent,
		FieldNameDocumentID,
		FieldNameDocumentName,
		FieldNameKBID,
		FieldNameChunkIndex,
	}
)
