package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

// Hiển thị thông tin category
func (r *Repo) GetCategory(ctx context.Context, category *model.Categories) error {
	// GetCategory là hàm lấy thông tin category
	if err := r.db.Where(category).First(category).Error; err != nil {
		return err
	}
	return nil
}

// Hiển thị thông tin tất cả category
func (r *Repo) GetAllCategory(ctx context.Context, category *[]model.Categories) error {
	// GetAllCategory là hàm lấy thông tin tất cả category
	if err := r.db.Find(category).Error; err != nil {
		return err
	}
	return nil
}

// Tạo mới category
func (r *Repo) CreateCategory(ctx context.Context, category *model.Categories) error {
	// CreateCategory là hàm tạo mới category
	// kiểm tra context đã bị hủy chưa
	select {
	case <-ctx.Done():
		return ctx.Err() // trả về lỗi context đã bị hủy
	default:
	}

	// thực hiện công việc tạo mới category
	if err := r.db.Create(category).Error; err != nil {
		return err
	}

	// nếu thành công thì trả về nil
	return nil
}

// Cập nhật thông tin của category
func (r *Repo) UpdateCategory(ctx context.Context, category *model.Categories) error {
	// UpdateCategory là hàm cập nhật thông tin category
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

// Xóa thông tin category nếu không có sản phẩm nào được đặt lịch hẹn
// Nhằm không bị ảnh hưởng đến doanh thu
func (r *Repo) DeleteCategory(ctx context.Context, category *model.Categories) error {
	// DeleteCategory là hàm xóa thông tin category
	if err := r.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
