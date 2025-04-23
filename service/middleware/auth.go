package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hildam/AI-Cloud-Drive/entity/errcode"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	AuthHeaderKey  = "Authorization"
	AuthBearerType = "Bearer"
)

// JWTAuth 中间件
func JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 获取 Authorization 头
			authHeader := c.Request().Header.Get(AuthHeaderKey)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    errcode.TokenMissing,
					"message": "需要认证令牌",
				})
			}

			// 分割 Bearer 和令牌
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != AuthBearerType {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    errcode.TokenInvalid,
					"message": "令牌格式错误",
				})
			}

			// 解析验证令牌
			tokenString := parts[1]
			claims, err := ParseToken(tokenString)
			if err != nil {
				status := http.StatusUnauthorized
				code := errcode.TokenInvalid
				message := "无效令牌"

				// 新版错误处理
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					message = "令牌已过期"
					code = errcode.TokenExpired
					status = http.StatusForbidden
				case errors.Is(err, jwt.ErrTokenMalformed):
					message = "令牌格式错误"
				case errors.Is(err, jwt.ErrSignatureInvalid):
					message = "签名验证失败"
				}

				return c.JSON(status, map[string]interface{}{
					"code":    code,
					"message": message,
				})
			}

			// 将用户ID存入上下文
			c.Set("user_id", claims.UserID)
			return next(c)
		}
	}
}
