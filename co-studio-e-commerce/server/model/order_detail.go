package model

import "github.com/google/uuid"

type OrderDetail struct {
	OrderDetailID int       `gorm:"primary_key;AUTO_INCREMENT;type:int" json:"order_detail_id"`
	OrderID       uuid.UUID `json:"order_id" gorm:"type:uuid"`
	ProductID     uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Quantity      int       `json:"quantity"`
	Subtotal      float64   `json:"subtotal"`
	Enable        int       `json:"enable"`
	IsUpdate      int       `json:"is_update"`
	WhoUpdate     string    `json:"who_update"`
	UpdateDate    string    `json:"update_date"`
	IsDelete      int       `json:"is_delete"`
	Order         Order     `gorm:"foreignKey:OrderID"`
	Product       Product   `gorm:"foreignKey:ProductID"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}

func (a *OrderDetail) IsSet() bool {
	return a.OrderDetailID != 0
}
