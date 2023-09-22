package model

import (
	"time"

	"github.com/google/uuid"
)

// activity_log tự động lấy thông tin hoạt động của người dùng
// thu thập hoạt động của người dùng
// thu thập lịch sử hoạt động của người dùng

type ActivityLog struct {
	Log_id        int       `gorm:"primary_key;AUTO_INCREMENT" json:"log_id"`
	User_id       uuid.UUID `json:"user_id"`
	Activity_type string    `json:"activity_type"`
	Activity_time time.Time `json:"activity_time"`
	Description   string    `json:"description"`
}

func (ActivityLog) TableName() string {
	return "activity_log"
}

func (a *ActivityLog) IsSet() bool {
	return a.Log_id != 0
}
