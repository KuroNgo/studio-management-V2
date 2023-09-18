package model

type Product struct {
	ProductID   int    `gorm:"primary_key;AUTO_INCREMENT" json:"product_id"`
	ProductName string `json:"product_name"`
	CategoryID  int    `json:"category_id"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image_URL   string `json:"image_url"`
	Enable      int    `json:"enable"`
	Is_Update   int    `json:"is_update"`
	Who_Update  string `json:"who_update"`
	Update_Date string `json:"update_date"`
	Is_Delete   int    `json:"is_delete"`
}

func (Product) TableName() string {
	return "product"
}

func (a *Product) IsSet() bool {
	return a.ProductID != 0
}
