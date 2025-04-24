package knowledge

//go:generate $GOPATH/bin/mockgen --source=knowledge.go --destination=knowledge_mock.go --package=knowledge

import (
	"context"

	"gorm.io/gorm"
)

type Dao interface {
	GetDB(ctx context.Context) *gorm.DB
	// 知识库相关
	CreateKB(ctx context.Context, kb *KnowledgeBase) error                                     // 创建知识库
	DeleteKB(ctx context.Context, id string) error                                             // 删除知识库
	CountKBs(ctx context.Context, userID uint) (int64, error)                                  // 统计知识库数量
	ListKBs(ctx context.Context, userID uint, page int, pageSize int) ([]KnowledgeBase, error) // 获取知识库列表
	GetKBByID(ctx context.Context, kb_id string) (*KnowledgeBase, error)                       // 获取知识库

	// 文档相关
	CreateDocument(ctx context.Context, doc *Document) error                         // 创建文档
	UpdateDocument(ctx context.Context, doc *Document) error                         // 更新文档
	CountDocs(ctx context.Context, id string) (int64, error)                         // 统计文档数量
	ListDocs(ctx context.Context, id string, page int, size int) ([]Document, error) // 获取文档列表
	GetAllDocsByKBID(ctx context.Context, kbID string) ([]Document, error)           // 获取知识库下所有文档
	DeleteDocsByKBID(ctx context.Context, kbID string) error                         // 删除知识库下所有文档
	BatchDeleteDocs(ctx context.Context, userID uint, docIDs []string) error         // 批量删除文档
}
