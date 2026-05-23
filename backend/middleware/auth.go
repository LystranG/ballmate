package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"ballmate/utils"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "未登录"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// 没有 Bearer 前缀
			c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "token格式错误"})
			c.Abort()
			return
		}

		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "token无效或已过期"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
