package model

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:category_id" json:"category_id"`
	CategoryName string    `gorm:"unique" json:"category_name"`
	Description  string    `json:"description"`
	Enable       int       `json:"enable" default:"1"`
	IsUpdate     int       `json:"is_update" default:"0"`
	WhoUpdate    string    `json:"who_update" default:"admin"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	IsDelete     int       `json:"is_delete" default:"0"`
}

type CategoryForCreate struct {
	CategoryName string `gorm:"unique" json:"category_name"`
	Description  string `json:"description"`
}

type CategoryForUpdate struct {
	CategoryName string `gorm:"unique" json:"category_name"`
	Description  string `json:"description"`
	Enable       int    `json:"enable" default:"1"`
}

type CategoryResponse struct {
	CategoryName string    `gorm:"unique" json:"category_name"`
	Description  string    `json:"description"`
	WhoUpdate    string    `json:"who_update" default:"admin"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (Category) TableName() string {
	return "categories"
}
