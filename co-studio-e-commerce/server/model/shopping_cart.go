package model

type ShoppingCart struct {
	ShoppingCartID int `gorm:"primary_key;AUTO_INCREMENT" json:"shopping_cart_id"`
	UserID         int `json:"user_id"`
	ProductID      int `json:"product_id"`
	Quantity       int `json:"quantity"`
	Enable         int `json:"enable"`
	Is_Update      int `json:"is_update"`
	Who_Update     int `json:"who_update"`
	Update_Date    int `json:"update_date"`
	Is_Delete      int `json:"is_delete"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}

func (a *ShoppingCart) IsSet() bool {
	return a.ShoppingCartID != 0
}
