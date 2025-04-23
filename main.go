package main

import (
	"log"

	"github.com/hildam/AI-Cloud-Drive/common/logger"
	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/hildam/AI-Cloud-Drive/logic/tracelog"
	"github.com/hildam/AI-Cloud-Drive/service"
	"github.com/hildam/AI-Cloud-Drive/service/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// 初始化配置
	if err := conf.Init(); err != nil {
		log.Fatalf("Failed to initialize config, err: %v", err)
	}

	// 初始化echo框架
	e := echo.New()
	e.Logger = logger.NewEchoLogger(&conf.GetCfg().Log) // 自定义日志
	e.HTTPErrorHandler = tracelog.EchoErrorHandler      // 自定义错误处理函数
	e.Use(getMiddlewares()...)

	// 配置路由
	service.SetRouter(e)

	// 启动服务器
	logger.Info().Msg("Starting server...")
	if err := e.Start(conf.GetCfg().Server.GetPort()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}

// getMiddlewares 中间件列表
func getMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		echomiddleware.Recover(), // 捕捉panic
		middleware.SetupCORS(),   // CORS处理
		tracelog.BodyDump(),      // 跟踪日志中间件
	}
}
