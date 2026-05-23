package handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"ballmate/service"
)

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Location        string `json:"location"`
	Sports          string `json:"sports"`
}

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册 handler
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请填写完整信息")
		return
	}

	// 校验密码一致性
	if req.Password != req.ConfirmPassword {
		Error(c, http.StatusBadRequest, "两次密码输入不一致")
		return
	}

	// 校验手机号格式（11位数字）
	phoneRegex := regexp.MustCompile(`^\d{11}$`)
	if !phoneRegex.MatchString(req.Phone) {
		Error(c, http.StatusBadRequest, "请输入正确的手机号")
		return
	}

	// 调用 service 注册
	user, err := service.Register(req.Username, req.Password, req.Phone, req.Location, req.Sports)
	if err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 注册成功后自动登录，获取 token
	token, _, err := service.Login(req.Username, req.Password)
	if err != nil {
		Error(c, http.StatusInternalServerError, "自动登录失败")
		return
	}

	Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
			"phone":    user.Phone,
			"avatar":   user.Avatar,
			"location": user.Location,
			"sports":   user.Sports,
		},
	})
}

// Login 用户登录 handler
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请填写账号和密码")
		return
	}

	token, user, err := service.Login(req.Username, req.Password)
	if err != nil {
		Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
			"phone":    user.Phone,
			"avatar":   user.Avatar,
			"location": user.Location,
			"sports":   user.Sports,
		},
	})
}
