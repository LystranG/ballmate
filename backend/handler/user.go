package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"ballmate/service"
)

// updateProfileRequest 更新资料请求体
type updateProfileRequest struct {
	Nickname string `json:"nickname"`
	Location string `json:"location"`
	Sports   string `json:"sports"`
}

// updatePasswordRequest 修改密码请求体
type updatePasswordRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// GetProfile 获取当前用户信息
func GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := service.GetProfile(userID)
	if err != nil {
		Error(c, http.StatusNotFound, err.Error())
		return
	}
	Success(c, user)
}

// UpdateProfile 更新用户资料
func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var req updateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	user, err := service.UpdateProfile(userID, req.Nickname, req.Location, req.Sports)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	Success(c, user)
}

// UpdatePassword 修改密码
func UpdatePassword(c *gin.Context) {
	userID := c.GetUint("userID")

	var req updatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 校验两次密码是否一致
	if req.NewPassword != req.ConfirmPassword {
		Error(c, http.StatusBadRequest, "两次密码输入不一致")
		return
	}

	if err := service.UpdatePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		Error(c, http.StatusBadRequest, err.Error())
		return
	}
	Success(c, nil)
}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	userID := c.GetUint("userID")

	file, err := c.FormFile("avatar")
	if err != nil {
		Error(c, http.StatusBadRequest, "请选择文件")
		return
	}

	// 校验文件大小（不超过 5MB）
	if file.Size > 5*1024*1024 {
		Error(c, http.StatusBadRequest, "图片大小不能超过5MB")
		return
	}

	// 校验文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !allowedExts[ext] {
		Error(c, http.StatusBadRequest, "请上传图片文件")
		return
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		Error(c, http.StatusInternalServerError, "文件保存失败")
		return
	}

	avatarPath := "/uploads/" + filename
	user, err := service.UpdateAvatar(userID, avatarPath)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	_ = user
	Success(c, gin.H{"avatar": avatarPath})
}
