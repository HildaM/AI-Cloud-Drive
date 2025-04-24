package dao

import (
	"context"
	"fmt"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/hildam/AI-Cloud-Drive/dao/milvus"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var milvusCli client.Client

// InitMilvus 初始化
func InitMilvus() (err error) {
	ctx := context.Background()
	// 初始化 cli
	milvusCli, err = client.NewClient(ctx, client.Config{
		Address: conf.GetCfg().Milvus.Address,
	})
	if err != nil {
		return fmt.Errorf("无法连接到Milvus: %w", err)
	}

	// 初始化文本chunks集合
	if err := initTextChunksCollection(ctx, milvusCli); err != nil {
		return fmt.Errorf("初始化集合失败: %w", err)
	}
	return nil
}

// GetMilVusCli 返回实例
func GetMilVusCli() client.Client {
	return milvusCli
}

func initTextChunksCollection(ctx context.Context, milvusClinet client.Client) error {
	// 从配置中获取集合名称
	milvusConfig := conf.GetCfg().Milvus
	collectionName := milvusConfig.CollectionName

	// 检查是否存在
	exists, err := milvusClinet.HasCollection(ctx, collectionName)
	if err != nil {
		return fmt.Errorf("检查集合存在失败: %w", err)
	}
	if exists {
		return nil
	}

	// 创建集合
	schema := &entity.Schema{
		CollectionName: collectionName,
		Description:    "存储文档分块和向量",
		AutoID:         false,
		Fields: []*entity.Field{
			{
				Name:       milvus.FieldNameID,
				DataType:   entity.FieldTypeVarChar,
				PrimaryKey: true,
				AutoID:     false,
				TypeParams: map[string]string{
					"max_length": milvusConfig.IDMaxLength,
				},
			},
			{
				Name:     milvus.FieldNameContent,
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": milvusConfig.ContentMaxLength,
				},
			},
			{
				Name:     milvus.FieldNameDocumentID,
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": milvusConfig.DocIDMaxLength,
				},
			},
			{
				Name:     milvus.FieldNameDocumentName,
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": milvusConfig.DocNameMaxLength,
				},
			},
			{
				Name:     milvus.FieldNameKBID,
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": milvusConfig.KbIDMaxLength,
				},
			},
			{
				Name:     milvus.FieldNameChunkIndex,
				DataType: entity.FieldTypeInt32,
			},
			{
				Name:     milvus.FieldNameVector,
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					"dim": fmt.Sprintf("%d", milvusConfig.VectorDimension),
				},
			},
		},
	}

	// 创建集合
	if err := milvusClinet.CreateCollection(ctx, schema, 1); err != nil {
		return fmt.Errorf("initTextChunksCollection failed, CreateCollection err: %+v", err)
	}

	// 构建索引
	idx, indexErr := milvusConfig.GetMilvusIndex()
	if indexErr != nil {
		return fmt.Errorf("initTextChunksCollection failed, GetMilvusIndex err: %+v", err)
	}

	// 创建索引
	if err := milvusClinet.CreateIndex(ctx, collectionName, "vector", idx, false); err != nil {
		return fmt.Errorf("initTextChunksCollection failed, CreateIndex err: %+v", err)
	}

	// 加载集合到内存
	if err := milvusClinet.LoadCollection(ctx, collectionName, false); err != nil {
		return fmt.Errorf("initTextChunksCollection failed, LoadCollection err: %+v", err)
	}
	return nil
}
