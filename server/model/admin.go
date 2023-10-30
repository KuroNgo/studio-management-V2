package model

import "github.com/google/uuid"

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:admin_id" json:"admin_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

// Đặt tên trong database là admin
func (Admin) TableName() string {
	return "admin"
}
