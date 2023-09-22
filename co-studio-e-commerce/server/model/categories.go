package model

type Categories struct {
	CategoryID   int    `gorm:"primary_key;AUTO_INCREMENT" json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Enable       int    `json:"enable" default:"1"`
	Is_Update    int    `json:"is_update" default:"0"`
	Who_Update   string `json:"who_update" default:"admin"`
	Update_Date  string `json:"update_date" default:"CURRENT_TIMESTAMP"`
	Is_Delete    int    `json:"is_delete" default:"0"`
}

func (Categories) TableName() string {
	return "categories"
}

func (a *Categories) IsSet() bool {
	return a.CategoryID != 0
}
