package model

import (
	"github.com/google/uuid"
	"time"
)

type ShoppingCart struct {
	ShoppingCartID int       `gorm:"primary_key;AUTO_INCREMENT" json:"shopping_cart_id"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid"`
	ProductID      uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Quantity       int       `json:"quantity"`
	Enable         int       `json:"enable"`
	IsUpdate       time.Time `json:"is_update"`
	WhoUpdate      time.Time `json:"who_update"`
	UpdateDate     time.Time `json:"update_date"`
	IsDelete       int       `json:"is_delete"`
	User           User      `gorm:"foreignKey:UserID"`
	Product        Product   `gorm:"foreignKey:ProductID"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}

func (a *ShoppingCart) IsSet() bool {
	return a.ShoppingCartID != 0
}
