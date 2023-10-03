package model

import "github.com/google/uuid"

type Categories struct {
	CategoryID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"category_id"`
	CategoryName string    `gorm:"unique" json:"category_name"`
	Description  string    `json:"description"`
	Enable       int       `json:"enable" default:"1"`
	Is_Update    int       `json:"is_update" default:"0"`
	Who_Update   string    `json:"who_update" default:"admin"`
	Update_Date  string    `json:"update_date" default:"CURRENT_TIMESTAMP"`
	Is_Delete    int       `json:"is_delete" default:"0"`
}

func (Categories) TableName() string {
	return "categories"
}
