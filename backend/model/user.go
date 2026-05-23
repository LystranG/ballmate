package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password string `gorm:"size:100;not null" json:"-"`
	Phone    string `gorm:"size:11;not null" json:"phone"`
	Nickname string `gorm:"size:50" json:"nickname"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Location string `gorm:"size:100" json:"location"`
	Sports   string `gorm:"size:255" json:"sports"`
}

// SetPassword 使用 bcrypt 加密密码
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword 验证密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
