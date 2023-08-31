package model

import "time"

type ActivityLog struct {
	log_id        int       `gorm:"primary_key;AUTO_INCREMENT" json:"log_id"`
	user_id       int       `json:"user_id"`
	activity_type string    `json:"activity_type"`
	activity_time time.Time `json:"activity_time"`
	description   string    `json:"description"`
}

func (ActivityLog) TableName() string {
	return "activity_log"
}

func (a *ActivityLog) IsSet() bool {
	return a.log_id != 0
}
