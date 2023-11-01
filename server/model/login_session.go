package model

import (
	"github.com/google/uuid"
	"time"
)

type LoginSession struct {
	SessionID  int       `gorm:"primary_key;AUTO_INCREMENT" json:"session_id"`
	UserID     uuid.UUID `json:"user_id"`
	Username   string    `json:"username"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTime time.Time `json:"logout_time"`
	User       User      `gorm:"foreignKey:UserID"`
}

func (LoginSession) TableName() string {
	return "login_session"
}
