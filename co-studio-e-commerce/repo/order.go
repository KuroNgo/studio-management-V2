package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

func (r *Repo) GetOrder(ctx context.Context, order *model.Order) error {
	// GetOrder là hàm lấy thông tin order
	if err := r.db.Where(order).First(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetAllOrder(ctx context.Context, order *[]model.Order) error {
	// GetAllOrder là hàm lấy thông tin tất cả order
	if err := r.db.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) CreateOrder(ctx context.Context, order *model.Order) error {
	// CreateOrder là hàm tạo mới order
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateOrder(ctx context.Context, order *model.Order) error {
	// UpdateOrder là hàm cập nhật thông tin order
	if err := r.db.Save(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteOrder(ctx context.Context, order *model.Order) error {
	// DeleteOrder là hàm xóa thông tin order
	if err := r.db.Delete(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) RemoveOrder(ctx context.Context, order *model.Order) error {
	// RemoveOrder là hàm xóa thông tin order
	// Unscoped() để xóa cả dữ liệu đã bị soft delete
	// soft delete là xóa dữ liệu nhưng không xóa vĩnh viễn, chỉ đánh dấu là đã xóa
	if err := r.db.Unscoped().Delete(order).Error; err != nil {
		return err
	}
	return nil
}
