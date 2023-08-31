package model

type Order struct {
	OrderID      int    `gorm:"primary_key;AUTO_INCREMENT" json:"order_id"`
	OrderDate    string `json:"order_date"`
	Total_Amount int    `json:"total_amount"`
	User_ID      int    `json:"user_id"`
	Enable       int    `json:"enable"`
	Is_Update    int    `json:"is_update"`
	Who_Update   string `json:"who_update"`
	Update_Date  string `json:"update_date"`
	Is_Delete    int    `json:"is_delete"`
}

func (Order) TableName() string {
	return "order"
}
