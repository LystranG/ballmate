package router

import (
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
		user.GET("/profile", handler.GetProfile)
		user.PUT("/profile", handler.UpdateProfile)
		user.PUT("/password", handler.UpdatePassword)
		user.POST("/avatar", handler.UploadAvatar)
	}

	// 邀约路由
	invitation := api.Group("/invitations")
	invitation.Use(middleware.AuthMiddleware())
	{
		invitation.POST("", handler.CreateInvitation)
		invitation.GET("/:id", handler.GetInvitation)
		invitation.PUT("/:id/terminate", handler.TerminateInvitation)
		invitation.DELETE("/:id", handler.DeleteInvitation)
		invitation.GET("/:id/participants", handler.GetParticipants)
	}

	// 静态文件服务（头像访问）
	r.Static("/uploads", "./uploads")
}
