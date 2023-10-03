package model

import (
	"github.com/google/uuid"
	"time"
)

type Reviews struct {
	ReviewID    int       `gorm:"primary_key;AUTO_INCREMENT" json:"review_id"`
	ProductID   uuid.UUID `json:"product_id"`
	UserID      uuid.UUID `json:"user_id"`
	ReviewText  string    `json:"review_text"`
	Rating      int       `json:"rating"`
	Review_Date time.Time `json:"review_date"`
	Enable      int       `json:"enable"`
	Is_Update   int       `json:"is_update"`
	Who_Update  string    `json:"who_update"`
	Update_Date string    `json:"update_date"`
	Is_Delete   int       `json:"is_delete"`
	User        User      `gorm:"foreignKey:UserID"`
	Product     Product   `gorm:"foreignKey:ProductID"`
}

func (Reviews) TableName() string {
	return "reviews"
}

func (a *Reviews) IsSet() bool {
	return a.ReviewID != 0
}
