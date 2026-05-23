package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ballmate/handler"
	"ballmate/middleware"
)

// SetupRouter 配置路由
func SetupRouter(r *gin.Engine) {
	api := r.Group("/api/v1")

	// 公开路由（无需认证）
	auth := api.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	// 需要认证的路由
	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", placeholderHandler)
		user.PUT("/profile", placeholderHandler)
		user.PUT("/password", placeholderHandler)
		user.POST("/avatar", placeholderHandler)
	}

	// 静态文件服务（头像访问）
	r.Static("/uploads", "./uploads")
}

// placeholderHandler 占位 handler，后续 plan 实现
func placeholderHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}
