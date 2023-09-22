package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

// Tạo mới category
func (r *Repo) CreateProduct(ctx context.Context, category *model.Categories) error {
	// CreateCategory là hàm tạo mới category
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}
