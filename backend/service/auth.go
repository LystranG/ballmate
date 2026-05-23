package service

import (
	"errors"

	"gorm.io/gorm"

	"ballmate/model"
	"ballmate/utils"
)

// DB 全局数据库实例，由 main.go 初始化后赋值
var DB *gorm.DB

// Register 用户注册
func Register(username, password, phone, location, sports string) (*model.User, error) {
	// 检查用户名是否已存在
	var existing model.User
	result := DB.Where("username = ?", username).First(&existing)
	if result.Error == nil {
		return nil, errors.New("该账号已被注册")
	}

	user := &model.User{
		Username: username,
		Phone:    phone,
		Nickname: username, // 默认昵称为用户名
		Location: location,
		Sports:   sports,
	}

	if err := user.SetPassword(password); err != nil {
		return nil, errors.New("密码加密失败")
	}

	if err := DB.Create(user).Error; err != nil {
		return nil, errors.New("注册失败，请稍后重试")
	}

	return user, nil
}

// Login 用户登录，返回 token
func Login(username, password string) (string, *model.User, error) {
	var user model.User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", nil, errors.New("账号或密码错误")
	}

	if !user.CheckPassword(password) {
		return "", nil, errors.New("账号或密码错误")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", nil, errors.New("token生成失败")
	}

	return token, &user, nil
}
