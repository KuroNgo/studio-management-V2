package model

import (
	"github.com/google/uuid"
	"time"
)

// activity_log tự động lấy thông tin hoạt động của người dùng
// thu thập hoạt động của người dùng
// thu thập lịch sử hoạt động của người dùng

type ActivityLog struct {
	LogID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"logid"`
	UserID       uuid.UUID `json:"user_id" `
	ActivityType string    `json:"activity_type"`
	ActivityTime time.Time `json:"activity_time"`
	Description  string    `json:"description"`
	User         User      `gorm:"foreignKey:UserID"`
}

func (ActivityLog) TableName() string {
	return "activity_log"
}

func (a *ActivityLog) IsSet() bool {
	return a.LogID != 0
}
