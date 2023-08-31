package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

func (r *Repo) GetCategory(ctx context.Context, category *model.Categories) error {
	// GetCategory là hàm lấy thông tin category
	if err := r.db.Where(category).First(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetAllCategory(ctx context.Context, category *[]model.Categories) error {
	// GetAllCategory là hàm lấy thông tin tất cả category
	if err := r.db.Find(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) CreateCategory(ctx context.Context, category *model.Categories) error {
	// CreateCategory là hàm tạo mới category
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateCategory(ctx context.Context, category *model.Categories) error {
	// UpdateCategory là hàm cập nhật thông tin category
	if err := r.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteCategory(ctx context.Context, category *model.Categories) error {
	// DeleteCategory là hàm xóa thông tin category
	if err := r.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
