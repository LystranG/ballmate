package model

import (
	"time"

	"gorm.io/gorm"
)

// Invitation 邀约模型
// Status 存储值: waiting(等待中) / gathered(已召集) / terminated(已终止)
// expired(已过期) 不存储，查询时根据 ActivityTime 动态计算
type Invitation struct {
	gorm.Model
	CreatorID    uint      `gorm:"not null;index" json:"creator_id"`
	SportType    string    `gorm:"size:20;not null" json:"sport_type"`
	ActivityTime time.Time `gorm:"not null" json:"activity_time"`
	Address      string    `gorm:"size:255;not null" json:"address"`
	Latitude     float64   `gorm:"not null" json:"latitude"`
	Longitude    float64   `gorm:"not null" json:"longitude"`
	MaxPeople    int       `gorm:"not null" json:"max_people"`
	Status       string    `gorm:"size:20;not null;default:waiting" json:"status"`
}
