package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"order_id"`
	OrderDate    time.Time `json:"order_date"`
	Total_Amount int32     `json:"total_amount"`
	UserID       uuid.UUID `json:"user_id" `
	Enable       int       `json:"enable"`
	Is_Update    int       `json:"is_update"`
	Who_Update   string    `json:"who_update"`
	Update_Date  string    `json:"update_date"`
	Is_Delete    int       `json:"is_delete"`
	User         User      `gorm:"foreignKey:UserID"`
}

func (Order) TableName() string {
	return "order"
}
