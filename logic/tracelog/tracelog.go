package tracelog

import (
	"github.com/hildam/AI-Cloud-Drive/common/ctxop"
	"github.com/hildam/AI-Cloud-Drive/common/logger"
	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BodyDump 跟踪日志echo中间件
func BodyDump() echo.MiddlewareFunc {
	skipper := func(c echo.Context) bool {
		// 提前添加公共字段，在dumpHandler中加的话，skipper到dump的中间过程日志无公共字段
		logger.WithFields(map[string]interface{}{
			"env":    conf.GetCfg().Server.Env,
			"method": c.Request().Method,
			"uri":    c.Request().URL.Path,
			"query":  c.Request().URL.RawQuery,
			"req":    string(ctxop.GetReqBody(c)),
		}).Msg("Request started")
		return false
	}
	dumpHandler := func(c echo.Context, req, rsp []byte) {
		logger.WithFields(map[string]interface{}{
			"rsp": string(rsp),
		}).Msg("Request completed")
	}
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{Skipper: skipper, Handler: dumpHandler})
}

// EchoErrorHandler echo框架自定义错误处理函数
func EchoErrorHandler(err error, c echo.Context) {
	c.Echo().DefaultHTTPErrorHandler(err, c)

	r := c.Request()
	logger.Error().
		Str("method", r.Method).
		Str("uri", r.RequestURI).
		Str("path", c.Path()).
		Err(err).
		Msg("Request error")
}
