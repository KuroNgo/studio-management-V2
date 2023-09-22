package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
)

type ActivityLog struct {
	service service.IActivityLog
}

type IActivityLog interface {
	GetActivityLogs()
}

func NewActivityLog(service service.IActivityLog) *ActivityLog {
	return &ActivityLog{service: service}
}

func (handler *ActivityLog) GetActivityLogs(ctx *gin.Context) {
	// get x-user-id from header
	userid, err := util.GetXUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	// get activity logs
	activityLogReq := model.ActivityLog{}
	if err := ctx.ShouldBindJSON(&activityLogReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}
	activityLogReq.User_id = userid

	// check valid request

	// // craete activity log
	// _activitylog, err := handler.service.Create(activityLogReq)
}
