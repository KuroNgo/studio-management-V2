package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:product_id" json:"product_id"`
	ProductName string    `gorm:"unique" json:"product_name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	AvatarURL   string    `json:"image_url"`
	Enable      int       `json:"enable"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeleteAt    gorm.DeletedAt
	IsUpdate    int       `json:"is_update"`
	IsDelete    int       `json:"is_delete"`
	WhoUpdates  string    `gorm:"not null;default:user" json:"who_updates"`
	CategoryID  uuid.UUID `json:"category_id" gorm:"type:uuid"`
	Categories  Category  `gorm:"foreignKey:CategoryID"`
}

type ProductForEdit struct {
	ProductName string    `gorm:"unique" json:"product_name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	AvatarURL   string    `json:"image_url"`
	CategoryID  uuid.UUID `json:"category_id" gorm:"type:uuid"`
}

type Image struct {
	ID          uuid.UUID
	Name        string
	ImageURL    string
	Description string
	Data        []byte
}

func (Product) TableName() string {
	return "product"
}
