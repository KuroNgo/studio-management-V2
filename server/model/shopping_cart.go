package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ShoppingCart struct {
	ShoppingCartID int       `gorm:"primary_key;AUTO_INCREMENT" json:"shopping_cart_id"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid"`
	ProductID      uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Quantity       int       `json:"quantity"`
	Enable         int       `json:"enable"`
	CreatedAt      time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt       gorm.DeletedAt
	IsUpdate       int     `json:"is_update"`
	IsDelete       int     `json:"is_delete"`
	WhoUpdates     string  `gorm:"not null;default:user" json:"who_updates"`
	User           User    `gorm:"foreignKey:UserID"`
	Product        Product `gorm:"foreignKey:ProductID"`
}

type ShoppingCartForEdit struct {
	Quantity int `json:"quantity"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
