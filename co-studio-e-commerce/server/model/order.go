package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:order_id" json:"order_id"`
	OrderDate   time.Time `json:"order_date"`
	TotalAmount int32     `json:"total_amount"`
	UserID      uuid.UUID `json:"user_id" `
	Enable      int       `json:"enable"`
	WhoUpdate   string    `json:"who_update"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (Order) TableName() string {
	return "order"
}
