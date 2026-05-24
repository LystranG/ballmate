package service

import (
	"errors"
	"time"

	"ballmate/model"
)

// InvitationResponse 邀约响应结构（含计算后状态和参与人数）
type InvitationResponse struct {
	ID               uint      `json:"id"`
	CreatorID        uint      `json:"creator_id"`
	SportType        string    `json:"sport_type"`
	ActivityTime     time.Time `json:"activity_time"`
	Address          string    `json:"address"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	MaxPeople        int       `json:"max_people"`
	Status           string    `json:"status"`
	ParticipantCount int       `json:"participant_count"`
	CreatedAt        time.Time `json:"created_at"`
}

// ParticipantInfo 参与人信息（手机号已脱敏）
type ParticipantInfo struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

// computeStatus 计算邀约实际状态
// waiting 状态下若活动时间已过则返回 expired
func computeStatus(inv model.Invitation) string {
	if inv.Status == "waiting" && inv.ActivityTime.Before(time.Now()) {
		return "expired"
	}
	return inv.Status
}

// maskPhone 手机号脱敏：138****5678
func maskPhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

func buildResponse(inv model.Invitation) *InvitationResponse {
	var count int64
	DB.Model(&model.Participation{}).Where("invitation_id = ?", inv.ID).Count(&count)
	return &InvitationResponse{
		ID:               inv.ID,
		CreatorID:        inv.CreatorID,
		SportType:        inv.SportType,
		ActivityTime:     inv.ActivityTime,
		Address:          inv.Address,
		Latitude:         inv.Latitude,
		Longitude:        inv.Longitude,
		MaxPeople:        inv.MaxPeople,
		Status:           computeStatus(inv),
		ParticipantCount: int(count),
		CreatedAt:        inv.CreatedAt,
	}
}

// CreateInvitation 创建邀约，创建者自动加入参与人
func CreateInvitation(creatorID uint, sportType, address string, activityTime time.Time, lat, lng float64, maxPeople int) (*InvitationResponse, error) {
	inv := model.Invitation{
		CreatorID:    creatorID,
		SportType:    sportType,
		ActivityTime: activityTime,
		Address:      address,
		Latitude:     lat,
		Longitude:    lng,
		MaxPeople:    maxPeople,
		Status:       "waiting",
	}
	if err := DB.Create(&inv).Error; err != nil {
		return nil, errors.New("创建邀约失败")
	}

	// 创建者自动加入参与人
	DB.Create(&model.Participation{InvitationID: inv.ID, UserID: creatorID})

	// 检查是否已满员（maxPeople=1 时创建即召集）
	var count int64
	DB.Model(&model.Participation{}).Where("invitation_id = ?", inv.ID).Count(&count)
	if int(count) >= maxPeople {
		DB.Model(&inv).Update("status", "gathered")
		inv.Status = "gathered"
	}

	return buildResponse(inv), nil
}

// GetInvitation 获取邀约详情
func GetInvitation(id uint) (*InvitationResponse, error) {
	var inv model.Invitation
	if err := DB.First(&inv, id).Error; err != nil {
		return nil, errors.New("邀约不存在")
	}
	return buildResponse(inv), nil
}

// TerminateInvitation 终止邀约（仅创建者、仅 waiting 状态）
func TerminateInvitation(id, userID uint) error {
	var inv model.Invitation
	if err := DB.First(&inv, id).Error; err != nil {
		return errors.New("邀约不存在")
	}
	if inv.CreatorID != userID {
		return errors.New("无权操作")
	}
	if computeStatus(inv) != "waiting" {
		return errors.New("只有等待中的邀约可以终止")
	}
	DB.Model(&inv).Update("status", "terminated")
	return nil
}

// DeleteInvitation 删除邀约（仅创建者、仅 terminated 状态）
func DeleteInvitation(id, userID uint) error {
	var inv model.Invitation
	if err := DB.First(&inv, id).Error; err != nil {
		return errors.New("邀约不存在")
	}
	if inv.CreatorID != userID {
		return errors.New("无权操作")
	}
	if inv.Status != "terminated" {
		return errors.New("只有已终止的邀约可以删除")
	}
	// 软删除邀约及其参与记录
	DB.Where("invitation_id = ?", id).Delete(&model.Participation{})
	DB.Delete(&inv)
	return nil
}

// GetParticipants 获取参与人列表（仅创建者可查看，手机号脱敏）
func GetParticipants(id, userID uint) ([]ParticipantInfo, error) {
	var inv model.Invitation
	if err := DB.First(&inv, id).Error; err != nil {
		return nil, errors.New("邀约不存在")
	}
	if inv.CreatorID != userID {
		return nil, errors.New("无权查看")
	}

	var results []struct {
		UserID   uint
		Nickname string
		Phone    string
	}
	DB.Table("participations").
		Select("participations.user_id, users.nickname, users.phone").
		Joins("JOIN users ON users.id = participations.user_id").
		Where("participations.invitation_id = ? AND participations.deleted_at IS NULL", id).
		Scan(&results)

	participants := make([]ParticipantInfo, len(results))
	for i, r := range results {
		participants[i] = ParticipantInfo{
			UserID:   r.UserID,
			Nickname: r.Nickname,
			Phone:    maskPhone(r.Phone),
		}
	}
	return participants, nil
}
