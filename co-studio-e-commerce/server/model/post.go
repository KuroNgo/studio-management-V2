package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"uniqueIndex;not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Image     string    `gorm:"not null" json:"image"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID"`
}

type CreatePostRequest struct {
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePost struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Image     string    `json:"image,omitempty"`
	User      string    `json:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
