package repo

import (
	"co-studio-e-commerce/model"
)

func (r *Repo) GetOrderDetail(orderDetail *model.OrderDetail) error {
	// GetOrderDetail là hàm lấy thông tin order_detail
	if err := r.db.Where(orderDetail).First(orderDetail).Error; err != nil {
		return err
	}
	return nil
}
