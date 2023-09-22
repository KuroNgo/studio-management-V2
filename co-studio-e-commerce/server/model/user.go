package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FullName  string    `gorm:"type:varchar(255);column:full_name" json:"fullName"`
	Username  string    `gorm:"type:varchar(255);not null;column:name" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null;unique;column:email" json:"email"`
	Password  string    `gorm:"not null;column:password" json:"password"`
	Phone     string    `gorm:"not null" json:"phone"`
	Role      string    `gorm:"type:varchar(255);not null;default:user" json:"role"`
	Provider  string    `gorm:"not null" json:"provider"`
	Photo     string    `gorm:"not null" json:"photo"`
	Verified  bool      `gorm:"not null" json:"verified"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
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
