package model

import (
	"github.com/google/uuid"
	"time"
)

type Login_Session struct {
	SessionID   int       `gorm:"primary_key;AUTO_INCREMENT" json:"session_id"`
	UserID      uuid.UUID `json:"user_id"`
	Username    string    `json:"username"`
	Login_Time  time.Time `json:"login_time"`
	Logout_Time time.Time `json:"logout_time"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (Login_Session) TableName() string {
	return "login_session"
}

func (a *Login_Session) IsSet() bool {
	return a.SessionID != 0
}
