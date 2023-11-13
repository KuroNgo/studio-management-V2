package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	// Tên thuộc tính được đặt trong golang phải là ID nếu kiểu dữ liệu là uuid.UUID
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:userid" json:"id"`
	FullName         string    `gorm:"type:varchar(255);column:full_name;" json:"fullName" binding:"required"`
	Username         string    `gorm:"type:varchar(255);column:username" json:"username" `
	Email            string    `gorm:"unique;column:email" json:"email" binding:"required"`
	Password         string    `gorm:"not null;column:password" json:"password" binding:"required"`
	Phone            string    `gorm:"column:phone" json:"phone" binding:"required"`
	Role             string    `gorm:"type:varchar(255);not null;default:user" json:"role"`
	Provider         string    `gorm:"not null;default:local" json:"provider"`
	AvatarUser       []byte    `json:"avatarUser"`
	Photo            []byte    `gorm:"column:photo; null" json:"photo"`
	VerificationCode string
	Verified         bool      `gorm:"not null;default:true" json:"verified"`
	CreatedAt        time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt         gorm.DeletedAt
	IsUpdate         int    `json:"is_update"`
	IsDelete         int    `json:"is_delete"`
	WhoUpdates       string `gorm:"not null;default:user" json:"who_updates"`
	Enable           int    `gorm:"not null;default:1" json:"enable"`
}

type SignInWithUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpInput struct {
	FullName string `gorm:"type:varchar(255);column:full_name" json:"fullName"`
	Username string `gorm:"type:varchar(255);column:username" json:"username"`
	Email    string `gorm:"unique;column:email" json:"email"`
	Password string `gorm:"not null;column:password" json:"password"`
	Phone    string `gorm:"column:phone" json:"phone"`
}

type SignInWithEmail struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     []byte    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Enable    int       `json:"enable"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
