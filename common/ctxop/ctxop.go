package ctxop

import (
	"context"
	"io"

	"github.com/labstack/echo/v4"
)

// GetCtx 获取原始请求上下文
func GetCtx(c echo.Context) context.Context {
	if c.Request() != nil {
		return c.Request().Context()
	}
	return context.Background()
}

// GetReqBody 读取请求body
func GetReqBody(c echo.Context) []byte {
	if c.Request().Body != nil {
		body, _ := io.ReadAll(c.Request().Body)
		return body
	}
	return nil
}
