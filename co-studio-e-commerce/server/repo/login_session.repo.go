package repo

import (
	"context"

	"co-studio-e-commerce/model"
)

func (r *Repo) ViewUserLoginTime(ctx context.Context, user *[]model.Login_Session) error {
	// ViewUserLogin là hàm lấy thông tin tất cả user login
	if err := r.db.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) ViewUserLogoutTime(ctx context.Context, user *[]model.Login_Session) error {
	// ViewUserLogout là hàm lấy thông tin tất cả user logout
	if err := r.db.Find(user).Error; err != nil {
		return err
	}
	return nil
}
