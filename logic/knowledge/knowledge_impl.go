package knowledge

import (
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
	"github.com/hildam/AI-Cloud-Drive/dao/milvus"
	"github.com/hildam/AI-Cloud-Drive/logic/file"
)

type knowledgeLogic struct {
	fileLogic      file.Logic      // 文件
	knowledgeDao   knowledge.Dao   // 知识库存储
	milvusDao      milvus.Dao      // 向量数据库存储
	embeddingLogic embedding.Logic // embedding 接口
}

func NewKnowledgeLogic() Logic {
	return &knowledgeLogic{
		fileLogic: file.NewFileLogic(),
	}
}
