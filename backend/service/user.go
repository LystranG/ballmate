package service

import (
	"errors"

	"ballmate/model"
)

// GetProfile 获取用户信息
func GetProfile(userID uint) (*model.User, error) {
	var user model.User
	if err := DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}

// UpdateProfile 更新用户资料（昵称、常住地、球类兴趣）
func UpdateProfile(userID uint, nickname, location, sports string) (*model.User, error) {
	var user model.User
	if err := DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("用户不存在")
	}

	DB.Model(&user).Updates(map[string]interface{}{
		"nickname": nickname,
		"location": location,
		"sports":   sports,
	})

	return &user, nil
}

// UpdatePassword 修改密码（验证旧密码后更新）
func UpdatePassword(userID uint, oldPassword, newPassword string) error {
	user, err := GetProfile(userID)
	if err != nil {
		return err
	}

	if !user.CheckPassword(oldPassword) {
		return errors.New("原密码错误")
	}

	if err := user.SetPassword(newPassword); err != nil {
		return errors.New("密码加密失败")
	}

	if err := DB.Save(user).Error; err != nil {
		return errors.New("密码更新失败")
	}

	return nil
}

// UpdateAvatar 更新头像路径
func UpdateAvatar(userID uint, avatarPath string) (*model.User, error) {
	var user model.User
	if err := DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("用户不存在")
	}

	DB.Model(&user).Updates(map[string]interface{}{
		"avatar": avatarPath,
	})

	return &user, nil
}
