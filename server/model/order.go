package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:order_id" json:"order_id"`
	OrderDate   time.Time `json:"order_date"`
	TotalAmount int32     `json:"total_amount"`
	UserID      uuid.UUID `json:"user_id" `
	Enable      int       `json:"enable"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt    gorm.DeletedAt
	IsUpdate    int    `json:"is_update"`
	IsDelete    int    `json:"is_delete"`
	WhoUpdates  string `gorm:"not null;default:user" json:"who_updates"`
	User        User   `gorm:"foreignKey:UserID"`
}

type OrderForEdit struct {
	OrderDate   time.Time `json:"order_date"`
	TotalAmount int32     `json:"total_amount"`
	UserID      uuid.UUID `json:"user_id" `
}

func (Order) TableName() string {
	return "order"
}
