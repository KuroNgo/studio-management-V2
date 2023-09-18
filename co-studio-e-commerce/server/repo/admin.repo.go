package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

// Hiển thị thông tin admin
func (r *Repo) GetAdmin(ctx context.Context, admin *model.Admin) error {
	// GetAdmin là hàm lấy thông tin admin
	if err := r.db.Where(admin).First(admin).Error; err != nil {
		return err
	}
	return nil
}

// Cập nhật thông tin của admin
func (r *Repo) UpdateAdmin(ctx context.Context, admin *model.Admin) error {
	// UpdateAdmin là hàm cập nhật thông tin admin
	if err := r.db.Save(admin).Error; err != nil {
		return err
	}
	return nil
}

//
