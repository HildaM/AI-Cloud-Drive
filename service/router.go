package service

import (
	"github.com/hildam/AI-Cloud-Drive/service/file"
	"github.com/hildam/AI-Cloud-Drive/service/knowledge"
	"github.com/hildam/AI-Cloud-Drive/service/middleware"
	"github.com/hildam/AI-Cloud-Drive/service/user"
	"github.com/labstack/echo/v4"
)

// SetRouter 设置路由
func SetRouter(e *echo.Echo) {
	g := e.Group("/api")

	// user 用户接口
	userGroup := g.Group("/users")
	userSvr := user.NewUserService()
	{
		userGroup.POST("/register", userSvr.Register)
		userGroup.POST("/login", userSvr.Login)
	}

	// 文件接口
	fileGroup := g.Group("/files")
	fileGroup.Use(middleware.JWTAuth()) // 启用认证
	fileSvr := file.NewFileService()
	{
		fileGroup.POST("/upload", fileSvr.Upload)
		fileGroup.GET("/page", fileSvr.PageList)
		fileGroup.GET("/download", fileSvr.Download)
		fileGroup.DELETE("/delete", fileSvr.Delete)
		fileGroup.POST("/folder", fileSvr.CreateFolder)
		fileGroup.POST("/move", fileSvr.BatchMove)
		fileGroup.GET("/search", fileSvr.Search)
		fileGroup.PUT("/rename", fileSvr.Rename)
		fileGroup.GET("/path", fileSvr.GetPath)
		fileGroup.GET("/idPath", fileSvr.GetIDPath)
	}

	// 知识库接口
	knowledgeGroup := g.Group("/knowledge")
	knowledgeGroup.Use(middleware.JWTAuth()) // 启用认证
	knowledgeSvr := knowledge.NewKnowledgeService()
	{
		knowledgeGroup.POST("/create", knowledgeSvr.Create)
		knowledgeGroup.DELETE("/delete", knowledgeSvr.Delete)
		knowledgeGroup.POST("/add", knowledgeSvr.AddExistFile)
		knowledgeGroup.POST("/addNew", knowledgeSvr.AddNewFile)
		knowledgeGroup.GET("/page", knowledgeSvr.PageList)
		knowledgeGroup.GET("/detail", knowledgeSvr.GetKBDetail)
		// Doc
		knowledgeGroup.GET("/docPage", knowledgeSvr.DocPage)
		knowledgeGroup.POST("/docDelete", knowledgeSvr.DeleteDocs)
		// RAG
		knowledgeGroup.POST("/retrieve", knowledgeSvr.Retrieve)
		knowledgeGroup.POST("/chat", knowledgeSvr.Chat)
		knowledgeGroup.POST("/stream", knowledgeSvr.ChatStream)
	}
}
