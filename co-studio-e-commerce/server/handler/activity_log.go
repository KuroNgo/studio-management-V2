package handler

import (
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
