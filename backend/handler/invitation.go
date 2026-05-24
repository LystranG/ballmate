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

// GetInvitation 获取邀约详情
func GetInvitation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的邀约ID")
		return
	}

	resp, err := service.GetInvitation(uint(id))
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
