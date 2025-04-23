package knowledge

import (
	"context"

	"github.com/hildam/AI-Cloud-Drive/dao/file"
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
)

type Logic interface {
	// 知识库
	CreateDB(name, description string, userID uint) error                               // 创建知识库
	DeleteKB(userID uint, kbID string) error                                            // 删除知识库
	PageList(userID uint, page int, size int) (int64, []knowledge.KnowledgeBase, error) // 获取知识库列表
	GetKBDetail(userID uint, kbID string) (*knowledge.KnowledgeBase, error)             // 获取知识库详情

	// 文档

	CreateDocument(userID uint, kbID string, file *file.File) (*knowledge.Document, error)     // 添加File到知识库
	ProcessDocument(doc *knowledge.Document) error                                             // 解析嵌入文档（后续需要细化）
	DocList(userID uint, kbID string, page int, size int) (int64, []knowledge.Document, error) // 获取知识库下的文件列表
	DeleteDocs(userID uint, docs []string) error                                               // 批量删除文件

	// RAG
	Retrieve(userID uint, kbID string, query string, topK int) ([]knowledge.Chunk, error)                              // 获取检索的Chunks
	RAGQuery(userID uint, query string, kbIDs []string) (*ChatResponse, error)                                         // 新增RAG查询方法
	RAGQueryStream(ctx context.Context, userID uint, query string, kbIDs []string) (<-chan *ChatStreamResponse, error) // 流式对话

	// TODO: 移动Document到其他知识库
	// TODO：修改知识库（名称、说明）
}
