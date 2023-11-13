package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:post_id" json:"post_id"`
	Title      string    `gorm:"uniqueIndex;not null" json:"title"`
	Content    string    `gorm:"not null" json:"content"`
	Image      string    `gorm:"not null" json:"image"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt   gorm.DeletedAt
	IsUpdate   int    `json:"is_update"`
	IsDelete   int    `json:"is_delete"`
	WhoUpdates string `gorm:"not null;default:user" json:"who_updates"`
}

type PostForEdit struct {
	Title   string `gorm:"uniqueIndex;not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	Image   string `gorm:"not null" json:"image"`
}
