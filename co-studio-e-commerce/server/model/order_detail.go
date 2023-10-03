package model

type OrderDetail struct {
	OrderDetailID int     `gorm:"primary_key;AUTO_INCREMENT" json:"order_detail_id"`
	OrderID       int     `json:"order_id" gorm:"type:int"`
	ProductID     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Subtotal      float64 `json:"subtotal"`
	Enable        int     `json:"enable"`
	Is_Update     int     `json:"is_update"`
	Who_Update    string  `json:"who_update"`
	Update_Date   string  `json:"update_date"`
	Is_Delete     int     `json:"is_delete"`
	Order         Order   `gorm:"foreignKey:OrderID"`
	Product       Product `gorm:"foreignKey:ProductID"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}

func (a *OrderDetail) IsSet() bool {
	return a.OrderDetailID != 0
}
