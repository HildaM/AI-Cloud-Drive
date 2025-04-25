package knowledge

import (
	"context"
	"gorm.io/gorm"
)

type knowledgeDao struct {
	db *gorm.DB // db实例
}

func NewKnowledgeDao(db *gorm.DB) Dao {
	return &knowledgeDao{db: db}
}

// 获取数据实例
func (k *knowledgeDao) GetDB(ctx context.Context) *gorm.DB

// 知识库相关

// 创建知识库
func (k *knowledgeDao) CreateKB(ctx context.Context, kb *KnowledgeBase) error

// 删除知识库
func (k *knowledgeDao) DeleteKB(ctx context.Context, id string) error

// 统计知识库数量
func (k *knowledgeDao) CountKBs(ctx context.Context, userID uint) (int64, error)

// 获取知识库列表
func (k *knowledgeDao) ListKBs(ctx context.Context, userID uint, page int, pageSize int) ([]KnowledgeBase, error)

// 获取知识库
func (k *knowledgeDao) GetKBByID(ctx context.Context, kb_id string) (*KnowledgeBase, error)

// 文档相关

// 创建文档
func (k *knowledgeDao) CreateDocument(ctx context.Context, doc *Document) error

// 更新文档
func (k *knowledgeDao) UpdateDocument(ctx context.Context, doc *Document) error

// 统计文档数量
func (k *knowledgeDao) CountDocs(ctx context.Context, id string) (int64, error)

// 获取文档列表
func (k *knowledgeDao) ListDocs(ctx context.Context, id string, page int, size int) ([]Document, error)

// 获取知识库下所有文档
func (k *knowledgeDao) GetAllDocsByKBID(ctx context.Context, kbID string) ([]Document, error)

// 删除知识库下所有文档
func (k *knowledgeDao) DeleteDocsByKBID(ctx context.Context, kbID string) error

// 批量删除文档
func (k *knowledgeDao) BatchDeleteDocs(ctx context.Context, userID uint, docIDs []string) error
