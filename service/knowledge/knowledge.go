package knowledge

import (
	"context"

	"github.com/hildam/AI-Cloud-Drive/logic/file"
	"github.com/hildam/AI-Cloud-Drive/logic/knowledge"
	"github.com/labstack/echo/v4"
)

type knowledgeService struct {
	knowledgeLogic knowledge.Logic
	fileLogic      file.Logic
}

func NewKnowledgeService(ctx context.Context) *knowledgeService {
	return &knowledgeService{
		knowledgeLogic: knowledge.NewKnowledgeLogic(),
		fileLogic:      file.NewFileLogic(ctx),
	}
}

func (s *knowledgeService) Create(c echo.Context) error {
	return nil
}

func (s *knowledgeService) Delete(c echo.Context) error {
	return nil
}

func (s *knowledgeService) AddExistFile(c echo.Context) error {
	return nil
}

func (s *knowledgeService) AddNewFile(c echo.Context) error {
	return nil
}

func (s *knowledgeService) PageList(c echo.Context) error {
	return nil
}

func (s *knowledgeService) GetKBDetail(c echo.Context) error {
	return nil
}

func (s *knowledgeService) DocPage(c echo.Context) error {
	return nil
}

func (s *knowledgeService) DeleteDocs(c echo.Context) error {
	return nil
}

func (s *knowledgeService) Retrieve(c echo.Context) error {
	return nil
}

func (s *knowledgeService) Chat(c echo.Context) error {
	return nil
}

func (s *knowledgeService) ChatStream(c echo.Context) error {
	return nil
}
