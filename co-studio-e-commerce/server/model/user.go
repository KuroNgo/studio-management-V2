package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primary_key;tableName:user" json:"id"  `
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Phone     string    `gorm:"not null" json:"phone"`
	Role      string    `gorm:"type:varchar(255);not null" json:"role"`
	Provider  string    `gorm:"not null" json:"provider"`
	Photo     string    `gorm:"not null" json:"photo"`
	Verified  bool      `gorm:"not null" json:"verified"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	Enable    int       `gorm:"not null" json:"enable"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Enable    int       `json:"enable"`
}
