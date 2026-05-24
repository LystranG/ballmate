package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ballmate/service"
)

// createInvitationRequest 创建邀约请求体
type createInvitationRequest struct {
	SportType    string  `json:"sport_type" binding:"required"`
	ActivityTime string  `json:"activity_time" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	Latitude     float64 `json:"latitude" binding:"required"`
	Longitude    float64 `json:"longitude" binding:"required"`
	MaxPeople    int     `json:"max_people" binding:"required,min=1"`
}

// CreateInvitation 创建邀约
func CreateInvitation(c *gin.Context) {
	userID := c.GetUint("userID")

	var req createInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	activityTime, err := time.Parse(time.RFC3339, req.ActivityTime)
	if err != nil {
		Error(c, http.StatusBadRequest, "时间格式错误，请使用 RFC3339 格式")
		return
	}

	resp, err := service.CreateInvitation(userID, req.SportType, req.Address, activityTime, req.Latitude, req.Longitude, req.MaxPeople)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	Success(c, resp)
}

// listNearbyRequest 附近邀约查询参数
type listNearbyRequest struct {
	Latitude  float64 `form:"lat" binding:"required"`
	Longitude float64 `form:"lng" binding:"required"`
	Page      int     `form:"page,default=1"`
	PageSize  int     `form:"page_size,default=10"`
	SortBy    string  `form:"sort_by,default=distance"` // distance | time
	SportType string  `form:"sport_type"`
}

// ListNearby 获取附近邀约列表（按距离或时间排序，支持球类筛选）
func ListNearby(c *gin.Context) {
	userID := c.GetUint("userID")
	var req listNearbyRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}
	// PageSize 上限 20，防止单次请求返回过多数据（T-03-02）
	if req.PageSize > 20 {
		req.PageSize = 20
	}

	list, total, err := service.ListNearbyInvitations(
		userID, req.Latitude, req.Longitude,
		req.SportType, req.SortBy, req.Page, req.PageSize,
	)
	if err != nil {
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	Success(c, gin.H{"list": list, "total": total})
}

// GetInvitation 获取邀约详情
func GetInvitation(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	resp, err := service.GetInvitation(uint(id), userID)
	if err != nil {
		Error(c, http.StatusNotFound, err.Error())
		return
	}
	Success(c, resp)
}

// TerminateInvitation 终止邀约
func TerminateInvitation(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	if err := service.TerminateInvitation(uint(id), userID); err != nil {
		Error(c, http.StatusForbidden, err.Error())
		return
	}
	Success(c, nil)
}

// DeleteInvitation 删除邀约
func DeleteInvitation(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	if err := service.DeleteInvitation(uint(id), userID); err != nil {
		Error(c, http.StatusForbidden, err.Error())
		return
	}
	Success(c, nil)
}

// GetParticipants 获取参与人列表
func GetParticipants(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	participants, err := service.GetParticipants(uint(id), userID)
	if err != nil {
		Error(c, http.StatusForbidden, err.Error())
		return
	}
	Success(c, participants)
}

// JoinInvitation 加入邀约
func JoinInvitation(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	if err := service.JoinInvitation(uint(id), userID); err != nil {
		msg := err.Error()
		// 权限类错误返回 403，业务类错误返回 400
		if msg == "您已加入该邀约" || msg == "邀约人数已满" || msg == "该邀约已无法加入" {
			Error(c, http.StatusBadRequest, msg)
		} else {
			Error(c, http.StatusForbidden, msg)
		}
		return
	}
	Success(c, nil)
}

// LeaveInvitation 退出邀约
func LeaveInvitation(c *gin.Context) {
	userID := c.GetUint("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	if err := service.LeaveInvitation(uint(id), userID); err != nil {
		msg := err.Error()
		if msg == "创建者不能退出自己的邀约" {
			Error(c, http.StatusForbidden, msg)
		} else {
			Error(c, http.StatusBadRequest, msg)
		}
		return
	}
	Success(c, nil)
}
