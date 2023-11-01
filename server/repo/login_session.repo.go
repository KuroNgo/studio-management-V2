package repo

import (
	"co-studio-e-commerce/model"
)

func (r *Repo) ViewUserLoginTime(user *[]model.LoginSession) error {
	// ViewUserLogin là hàm lấy thông tin tất cả user login
	if err := r.db.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) ViewUserLogoutTime(user *[]model.LoginSession) error {
	// ViewUserLogout là hàm lấy thông tin tất cả user logout
	if err := r.db.Find(user).Error; err != nil {
		return err
	}
	return nil
}
