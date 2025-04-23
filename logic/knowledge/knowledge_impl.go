package knowledge

import (
	"github.com/hildam/AI-Cloud-Drive/dao/file"
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
)

type knowledgeLogic struct {
	fileLogic file.Logic
}

func NewKnowledgeLogic() Logic {
	return &knowledgeLogic{
		fileLogic: file.NewFileLogic(),
	}
}
