package model

import "gorm.io/gorm"

// Participation 参与记录模型
// 联合唯一索引防止同一用户重复参与同一邀约
type Participation struct {
	gorm.Model
	InvitationID uint `gorm:"not null;index;uniqueIndex:idx_inv_user" json:"invitation_id"`
	UserID       uint `gorm:"not null;index;uniqueIndex:idx_inv_user" json:"user_id"`
}
