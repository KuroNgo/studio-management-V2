package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt   gorm.DeletedAt
	WhoUpdates string `gorm:"not null;default:user" json:"who_updates"`
}
