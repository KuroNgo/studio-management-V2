package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type OrderDetail struct {
	OrderDetailID int       `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"order_detail_id"`
	OrderID       uuid.UUID `json:"order_id" gorm:"type:uuid"`
	ProductID     uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Quantity      int       `json:"quantity"`
	Subtotal      float64   `json:"subtotal"`
	Enable        int       `json:"enable"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt      gorm.DeletedAt
	IsUpdate      int     `json:"is_update"`
	IsDelete      int     `json:"is_delete"`
	WhoUpdates    string  `gorm:"not null;default:user" json:"who_updates"`
	Order         Order   `gorm:"foreignKey:OrderID"`
	Product       Product `gorm:"foreignKey:ProductID"`
}

type OrderDetailForEdit struct {
	OrderID   uuid.UUID `json:"order_id" gorm:"type:uuid"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
