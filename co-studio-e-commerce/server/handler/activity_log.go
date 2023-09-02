package handler

import (
	"github.com/gin-gonic/gin"

	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
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

// hiển thị hoạt động
// chỉ có user đăng nhập mới thấy
func (handler *ActivityLog) GetActivityLogs(ctx *gin.Context, oneActivityLog model.ActivityLog) (activityLogRes model.ActivityLog, err error) {

}
