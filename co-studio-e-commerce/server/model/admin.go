package model

import "github.com/google/uuid"

type Admin struct {
	AdminID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"adminID"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
}

// Đặt tên trong database là admin
func (Admin) TableName() string {
	return "admin"
}