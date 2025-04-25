package knowledge

import (
	"context"

	"github.com/hildam/AI-Cloud-Drive/dao/file"
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
	"github.com/hildam/AI-Cloud-Drive/dao/milvus"
)

type Logic interface {
	// 知识库

	// 创建知识库
	CreateDB(ctx context.Context, name, description string, userID uint) error

	// 删除知识库
	DeleteKB(ctx context.Context, userID uint, kbID string) error

	// 获取知识库列表
	PageList(ctx context.Context, userID uint, page int, size int) (int64, []knowledge.KnowledgeBase, error)

	// 获取知识库详情
	GetKBDetail(ctx context.Context, userID uint, kbID string) (*knowledge.KnowledgeBase, error)

	// 文档

	// 添加File到知识库
	CreateDocument(ctx context.Context, userID uint, kbID string, file *file.File) (*knowledge.Document, error)

	// 解析嵌入文档（后续需要细化）
	ProcessDocument(ctx context.Context, doc *knowledge.Document) error

	// 获取知识库下的文件列表
	DocList(ctx context.Context, userID uint, kbID string, page int, size int) (int64, []knowledge.Document, error)

	// 批量删除文件
	DeleteDocs(ctx context.Context, userID uint, docs []string) error

	// RAG

	// 获取检索的Chunks
	Retrieve(ctx context.Context, userID uint, kbID string, query string, topK int) ([]milvus.Chunk, error)

	// 新增RAG查询方法
	RAGQuery(ctx context.Context, userID uint, query string, kbIDs []string) (*ChatResponse, error)

	// 流式对话
	RAGQueryStream(ctx context.Context, userID uint, query string, kbIDs []string) (<-chan *ChatStreamResponse, error)

	// TODO: 移动Document到其他知识库
	// TODO：修改知识库（名称、说明）
}
