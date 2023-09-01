package model

type Categories struct {
	CategoryID   int    `gorm:"primary_key;AUTO_INCREMENT" json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Enable       int    `json:"enable"`
	Is_Update    int    `json:"is_update"`
	Who_Update   string `json:"who_update"`
	Update_Date  string `json:"update_date"`
	Is_Delete    int    `json:"is_delete"`
}

func (Categories) TableName() string {
	return "categories"
}

func (a *Categories) IsSet() bool {
	return a.CategoryID != 0
}
