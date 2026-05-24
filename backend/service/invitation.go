package service

import (
	"errors"
	"math"
	"sort"
	"time"

	"ballmate/model"
)

const earthRadiusKm = 6371.0

// haversineDistance 计算两点间球面距离（公里），作为 SQL 不可用时的回退
func haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c
}

// NearbyInvitationResponse 附近邀约响应（含距离字段）
type NearbyInvitationResponse struct {
	InvitationResponse
	Distance float64 `json:"distance"`
}

// ListNearbyInvitations 查询附近邀约（Haversine 距离计算 + 排序 + 筛选 + 分页）
// 排除当前用户自己创建的邀约（per NEARBY-02）
func ListNearbyInvitations(userID uint, lat, lng float64, sportType, sortBy string, page, pageSize int) ([]NearbyInvitationResponse, int64, error) {
	// Haversine SQL 表达式（SQLite 不支持 RADIANS，手动 * pi/180）
	// mattn/go-sqlite3 默认启用 SQLITE_ENABLE_MATH_FUNCTIONS，支持 sin/cos/asin/sqrt/pow
	distanceExpr := `(6371 * 2 * asin(sqrt(
		pow(sin((latitude - ?) * 3.141592653589793 / 180 / 2), 2) +
		cos(? * 3.141592653589793 / 180) * cos(latitude * 3.141592653589793 / 180) *
		pow(sin((longitude - ?) * 3.141592653589793 / 180 / 2), 2)
	)))`

	query := DB.Table("invitations").
		Select("invitations.*, "+distanceExpr+" as distance", lat, lat, lng).
		Where("invitations.deleted_at IS NULL").
		Where("invitations.creator_id != ?", userID)

	// 球类筛选
	if sportType != "" {
		query = query.Where("invitations.sport_type = ?", sportType)
	}

	// 总数统计（不含分页）
	countQuery := DB.Table("invitations").
		Where("deleted_at IS NULL").
		Where("creator_id != ?", userID)
	if sportType != "" {
		countQuery = countQuery.Where("sport_type = ?", sportType)
	}
	var total int64
	countQuery.Count(&total)

	// 排序
	if sortBy == "time" {
		query = query.Order("activity_time ASC")
	} else {
		query = query.Order("distance ASC")
	}

	// 分页
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	var results []struct {
		model.Invitation
		Distance float64
	}
	if err := query.Scan(&results).Error; err != nil {
		// SQL 数学函数不可用时回退到 Go 层计算
		return listNearbyFallback(userID, lat, lng, sportType, sortBy, page, pageSize)
	}

	// 检查是否因数学函数不可用导致 distance 全为 0（非正常情况）
	// 若所有 distance 为 0 且有多条记录，说明 SQL 计算失败，回退到 Go 层
	if len(results) > 1 {
		allZero := true
		for _, r := range results {
			if r.Distance != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return listNearbyFallback(userID, lat, lng, sportType, sortBy, page, pageSize)
		}
	}

	responses := make([]NearbyInvitationResponse, len(results))
	for i, r := range results {
		base := buildResponse(r.Invitation)
		responses[i] = NearbyInvitationResponse{
			InvitationResponse: *base,
			Distance:           r.Distance,
		}
	}
	return responses, total, nil
}

// listNearbyFallback 当 SQLite 数学函数不可用时，Go 层计算距离并排序分页
func listNearbyFallback(userID uint, lat, lng float64, sportType, sortBy string, page, pageSize int) ([]NearbyInvitationResponse, int64, error) {
	query := DB.Table("invitations").
		Where("deleted_at IS NULL").
		Where("creator_id != ?", userID)
	if sportType != "" {
		query = query.Where("sport_type = ?", sportType)
	}

	var invitations []model.Invitation
	query.Find(&invitations)

	// Go 层计算距离
	type withDist struct {
		inv      model.Invitation
		distance float64
	}
	items := make([]withDist, len(invitations))
	for i, inv := range invitations {
		items[i] = withDist{inv: inv, distance: haversineDistance(lat, lng, inv.Latitude, inv.Longitude)}
	}

	// 排序
	if sortBy == "time" {
		sort.Slice(items, func(i, j int) bool {
			return items[i].inv.ActivityTime.Before(items[j].inv.ActivityTime)
		})
	} else {
		sort.Slice(items, func(i, j int) bool {
			return items[i].distance < items[j].distance
		})
	}

	total := int64(len(items))

	// 分页
	offset := (page - 1) * pageSize
	if offset >= len(items) {
		return []NearbyInvitationResponse{}, total, nil
	}
	end := offset + pageSize
	if end > len(items) {
		end = len(items)
	}
	items = items[offset:end]

	responses := make([]NearbyInvitationResponse, len(items))
	for i, item := range items {
		base := buildResponse(item.inv)
		responses[i] = NearbyInvitationResponse{
			InvitationResponse: *base,
			Distance:           item.distance,
		}
	}
	return responses, total, nil
}

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
	HasJoined        bool      `json:"has_joined"`
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

// MyInvitations 查询当前用户的邀约列表
// listType == "created"：查询用户发起的邀约
// listType == "joined"：查询用户参与但非自己创建的邀约
func MyInvitations(userID uint, listType string, page, pageSize int) ([]InvitationResponse, int64, error) {
	offset := (page - 1) * pageSize

	var invitations []model.Invitation
	var total int64

	if listType == "joined" {
		// 查询用户参与但非自己创建的邀约（排除自己创建的）
		DB.Table("invitations").
			Joins("JOIN participations ON participations.invitation_id = invitations.id").
			Where("participations.user_id = ? AND participations.deleted_at IS NULL AND invitations.creator_id != ? AND invitations.deleted_at IS NULL", userID, userID).
			Count(&total)

		DB.Table("invitations").
			Joins("JOIN participations ON participations.invitation_id = invitations.id").
			Where("participations.user_id = ? AND participations.deleted_at IS NULL AND invitations.creator_id != ? AND invitations.deleted_at IS NULL", userID, userID).
			Order("invitations.created_at DESC").
			Offset(offset).Limit(pageSize).
			Find(&invitations)
	} else {
		// 默认 created：查询用户发起的邀约
		DB.Model(&model.Invitation{}).
			Where("creator_id = ? AND deleted_at IS NULL", userID).
			Count(&total)

		DB.Model(&model.Invitation{}).
			Where("creator_id = ? AND deleted_at IS NULL", userID).
			Order("created_at DESC").
			Offset(offset).Limit(pageSize).
			Find(&invitations)
	}

	responses := make([]InvitationResponse, len(invitations))
	for i, inv := range invitations {
		responses[i] = *buildResponse(inv)
	}
	return responses, total, nil
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

// GetInvitation 获取邀约详情（含当前用户是否已加入）
func GetInvitation(id, userID uint) (*InvitationResponse, error) {
	var inv model.Invitation
	if err := DB.First(&inv, id).Error; err != nil {
		return nil, errors.New("邀约不存在")
	}
	resp := buildResponse(inv)
	// 查询当前用户是否已加入
	var joinCount int64
	DB.Model(&model.Participation{}).
		Where("invitation_id = ? AND user_id = ? AND deleted_at IS NULL", id, userID).
		Count(&joinCount)
	resp.HasJoined = joinCount > 0
	return resp, nil
}

// JoinInvitation 加入邀约
// 校验：邀约存在、状态为 waiting、未满员、未重复加入
// 加入后若满员则自动将状态流转为 gathered
func JoinInvitation(invitationID, userID uint) error {
	var inv model.Invitation
	if err := DB.First(&inv, invitationID).Error; err != nil {
		return errors.New("邀约不存在")
	}
	if computeStatus(inv) != "waiting" {
		return errors.New("该邀约已无法加入")
	}
	// 校验未满员
	var count int64
	DB.Model(&model.Participation{}).Where("invitation_id = ? AND deleted_at IS NULL", invitationID).Count(&count)
	if int(count) >= inv.MaxPeople {
		return errors.New("邀约人数已满")
	}
	// 创建参与记录（联合唯一索引自动防重复）
	if err := DB.Create(&model.Participation{InvitationID: invitationID, UserID: userID}).Error; err != nil {
		return errors.New("您已加入该邀约")
	}
	// 加入后重新统计，满员则自动流转为 gathered
	var newCount int64
	DB.Model(&model.Participation{}).Where("invitation_id = ? AND deleted_at IS NULL", invitationID).Count(&newCount)
	if int(newCount) >= inv.MaxPeople {
		DB.Model(&inv).Update("status", "gathered")
	}
	return nil
}

// LeaveInvitation 退出邀约
// 校验：邀约存在、非创建者、状态为 waiting、已加入
// 退出后若邀约之前是 gathered 则恢复为 waiting
func LeaveInvitation(invitationID, userID uint) error {
	var inv model.Invitation
	if err := DB.First(&inv, invitationID).Error; err != nil {
		return errors.New("邀约不存在")
	}
	if inv.CreatorID == userID {
		return errors.New("创建者不能退出自己的邀约")
	}
	if computeStatus(inv) != "waiting" {
		return errors.New("该邀约状态不允许退出")
	}
	// 软删除参与记录
	result := DB.Where("invitation_id = ? AND user_id = ?", invitationID, userID).Delete(&model.Participation{})
	if result.RowsAffected == 0 {
		return errors.New("您未加入该邀约")
	}
	// 若邀约之前是 gathered（满员），退出后人数不足，恢复为 waiting
	if inv.Status == "gathered" {
		DB.Model(&inv).Update("status", "waiting")
	}
	return nil
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
