package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

func (r *Repo) GetOrderDetail(ctx context.Context, order_detail *model.OrderDetail) error {
	// GetOrderDetail là hàm lấy thông tin order_detail
	if err := r.db.Where(order_detail).First(order_detail).Error; err != nil {
		return err
	}
	return nil
}
