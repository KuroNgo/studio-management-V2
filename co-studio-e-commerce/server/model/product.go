package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"product_id"`
	ProductName string     `gorm:"unique" json:"product_name"`
	CategoryID  int        `json:"category_id" gorm:"type:int"`
	Price       int        `json:"price"`
	Description string     `json:"description"`
	Image_URL   string     `json:"image_url"`
	Enable      int        `json:"enable"`
	Is_Update   int        `json:"is_update"`
	Who_Update  string     `json:"who_update"`
	Update_Date string     `json:"update_date"`
	Is_Delete   int        `json:"is_delete"`
	Categories  Categories `gorm:"foreignKey:CategoryID"`
}

type Image struct {
	gorm.Model
	Name        string
	Description string
	Data        []byte
}

func (Product) TableName() string {
	return "product"
}
