package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

func (r *Repo) GetAdmin(ctx context.Context, admin *model.Admin) error {
	// GetAdmin là hàm lấy thông tin admin
	if err := r.db.Where(admin).First(admin).Error; err != nil {
		return err
	}
	return nil
}
