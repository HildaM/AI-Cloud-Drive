package middleware

import (
	"time"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupCORS 封装CORS配置
func SetupCORS() echo.MiddlewareFunc {
	corsConfig := conf.GetCfg().CORS

	maxAge, err := time.ParseDuration(corsConfig.MaxAge)
	if err != nil {
		maxAge = 12 * time.Hour
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     corsConfig.AllowOrigins,     // 允许所有域名
		AllowMethods:     corsConfig.AllowMethods,     // 允许的HTTP方法
		AllowHeaders:     corsConfig.AllowHeaders,     // 允许的请求头
		ExposeHeaders:    corsConfig.ExposeHeaders,    // 暴露的响应头
		AllowCredentials: corsConfig.AllowCredentials, // 允许携带凭证（如Cookie）
		MaxAge:           int(maxAge.Seconds()),       // 预检请求缓存时间
	})
}
