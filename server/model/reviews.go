package model

import (
	"github.com/google/uuid"
	"time"
)

type Reviews struct {
	ReviewID   int       `gorm:"primary_key;AUTO_INCREMENT" json:"review_id"`
	ProductID  uuid.UUID `json:"product_id" gorm:"type:uuid"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
	ReviewText string    `json:"review_text"`
	Rating     int       `json:"rating"`
	ReviewDate time.Time `json:"review_date"`
	Enable     int       `json:"enable"`
	IsUpdate   int       `json:"is_update"`
	WhoUpdate  string    `json:"who_update"`
	UpdateDate time.Time `json:"update_date"`
	IsDelete   int       `json:"is_delete"`
	User       User      `gorm:"foreignKey:UserID"`
	Product    Product   `gorm:"foreignKey:ProductID"`
}

func (Reviews) TableName() string {
	return "reviews"
}
