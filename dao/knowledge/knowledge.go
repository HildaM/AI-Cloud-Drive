package knowledge

//go:generate $GOPATH/bin/mockgen --source=knowledge.go --destination=knowledge_mock.go --package=knowledge

import (
	"context"

	"gorm.io/gorm"
)

type Dao interface {
	// 获取数据实例
	GetDB(ctx context.Context) *gorm.DB

	// 知识库相关

	// 创建知识库
	CreateKB(ctx context.Context, kb *KnowledgeBase) error

	// 删除知识库
	DeleteKB(ctx context.Context, id string) error

	// 统计知识库数量
	CountKBs(ctx context.Context, userID uint) (int64, error)

	// 获取知识库列表
	ListKBs(ctx context.Context, userID uint, page int, pageSize int) ([]KnowledgeBase, error)

	// 获取知识库
	GetKBByID(ctx context.Context, kb_id string) (*KnowledgeBase, error)

	// 文档相关

	// 创建文档
	CreateDocument(ctx context.Context, doc *Document) error

	// 更新文档
	UpdateDocument(ctx context.Context, doc *Document) error

	// 统计文档数量
	CountDocs(ctx context.Context, id string) (int64, error)

	// 获取文档列表
	ListDocs(ctx context.Context, id string, page int, size int) ([]Document, error)

	// 获取知识库下所有文档
	GetAllDocsByKBID(ctx context.Context, kbID string) ([]Document, error)

	// 删除知识库下所有文档
	DeleteDocsByKBID(ctx context.Context, kbID string) error

	// 批量删除文档
	BatchDeleteDocs(ctx context.Context, userID uint, docIDs []string) error
}
