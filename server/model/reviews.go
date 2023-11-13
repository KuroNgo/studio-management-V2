package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt   gorm.DeletedAt
	IsUpdate   int     `json:"is_update"`
	IsDelete   int     `json:"is_delete"`
	WhoUpdates string  `gorm:"not null;default:user" json:"who_updates"`
	User       User    `gorm:"foreignKey:UserID"`
	Product    Product `gorm:"foreignKey:ProductID"`
}

type ReviewForEdit struct {
	ReviewText string    `json:"review_text"`
	Rating     int       `json:"rating"`
	ReviewDate time.Time `json:"review_date"`
}

type ReviewResponse struct {
	ReviewText string    `json:"review_text"`
	Rating     int       `json:"rating"`
	ReviewDate time.Time `json:"review_date"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	ProductID  uuid.UUID `json:"product_id" gorm:"type:uuid"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

func (Reviews) TableName() string {
	return "reviews"
}
