package knowledge

import (
	"context"
	"log"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/hildam/AI-Cloud-Drive/dao"
	fileDao "github.com/hildam/AI-Cloud-Drive/dao/file"
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
	"github.com/hildam/AI-Cloud-Drive/dao/milvus"
	"github.com/hildam/AI-Cloud-Drive/logic/embedding"
	"github.com/hildam/AI-Cloud-Drive/logic/file"
)

type knowledgeLogic struct {
	fileLogic      file.Logic      // 文件
	knowledgeDao   knowledge.Dao   // 知识库存储
	milvusDao      milvus.Dao      // 向量数据库存储
	embeddingLogic embedding.Logic // embedding 接口
}

// NewKnowledgeLogic 创建实例
func NewKnowledgeLogic(ctx context.Context) Logic {
	embed, err := embedding.NewEmbedding(ctx, &conf.GetCfg().Embedding)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &knowledgeLogic{
		fileLogic:      file.NewFileLogic(ctx),
		knowledgeDao:   knowledge.NewKnowledgeDao(dao.GetDb()),
		milvusDao:      milvus.NewMilvusDao(dao.GetMilVusCli()),
		embeddingLogic: embed,
	}
}

// 创建知识库
func (k *knowledgeLogic) CreateDB(ctx context.Context, name, description string, userID uint) error

// 删除知识库
func (k *knowledgeLogic) DeleteKB(ctx context.Context, userID uint, kbID string) error

// 获取知识库列表
func (k *knowledgeLogic) PageList(ctx context.Context, userID uint, page int, size int) (int64, []knowledge.KnowledgeBase, error)

// 获取知识库详情
func (k *knowledgeLogic) GetKBDetail(ctx context.Context, userID uint, kbID string) (*knowledge.KnowledgeBase, error)

// 添加File到知识库
func (k *knowledgeLogic) CreateDocument(ctx context.Context, userID uint, kbID string, file *fileDao.File) (*knowledge.Document, error)

// 解析嵌入文档（后续需要细化）
func (k *knowledgeLogic) ProcessDocument(ctx context.Context, doc *knowledge.Document) error

// 获取知识库下的文件列表
func (k *knowledgeLogic) DocList(ctx context.Context, userID uint, kbID string, page int, size int) (int64, []knowledge.Document, error)

// 批量删除文件
func (k *knowledgeLogic) DeleteDocs(ctx context.Context, userID uint, docs []string) error

// 获取检索的Chunks
func (k *knowledgeLogic) Retrieve(ctx context.Context, userID uint, kbID string, query string, topK int) ([]milvus.Chunk, error)

// 新增RAG查询方法
func (k *knowledgeLogic) RAGQuery(ctx context.Context, userID uint, query string, kbIDs []string) (*ChatResponse, error)

// 流式对话
func (k *knowledgeLogic) RAGQueryStream(ctx context.Context, userID uint, query string, kbIDs []string) (<-chan *ChatStreamResponse, error)
