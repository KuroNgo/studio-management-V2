package repo

import (
	"co-studio-e-commerce/model"
)

func (r *Repo) GetUser(user *model.User) error {
	// GetUser là hàm lấy thông tin user
	if err := r.db.Where(user).First(user).Error; err != nil {
		return err
	}
	return nil
}

// Done
func (r *Repo) GetAllUser() ([]model.User, error) {
	// GetAllUser là hàm lấy thông tin tất cả user
	var users []model.User

	// Thực hiện truy vấn hoặc lấy thông tin tất cả người dùng từ nguồn dữ liệu của bạn
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err // Trả về nil và lỗi nếu có lỗi
	}

	return users, nil // Trả về danh sách người dùng và không có lỗi nếu thành công
}

func (r *Repo) CreateUser(user model.User) (model.User, error) {
	// CreateUser là hàm tạo mới user

	// Thực hiện tạo mới user trong cơ sở dữ liệu
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err // Trả về lỗi nếu có lỗi
	}

	return user, nil // Trả về thông tin người dùng và không có lỗi nếu thành công
}

func (r *Repo) UpdateUser(user model.User) (model.User, error) {
	// UpdateUser là hàm cập nhật thông tin user
	r.db.Save(user)
	return user, nil
}

// lập lịch xóa ( admin không có quyền xóa user )
func (r *Repo) DeleteUser(user model.User) (model.User, error) {
	// DeleteUser là hàm xóa thông tin user
	r.db.Delete(user)
	return user, nil
}

// thay đổi trạng thái user
// thay đổi trạng thái có 2 tác động: admin hoặc tự động
// admin: thay đổi trạng thái của user
// tự động: thay đổi trạng thái của user khi hết hạn ( hoặc user bị dính black list)
func (r *Repo) ChangeUserStatus(user *model.User) error {
	// ChangeUserStatus là hàm thay đổi trạng thái user
	if err := r.db.Model(user).Update("status", user.Enable).Error; err != nil {
		return err
	}
	return nil
}

// get id user
// login
func (r *Repo) GetUserID(id int) (model.User, error) {
	// GetUserID là hàm lấy thông tin user
	var user model.User
	r.db.Where("id = ?", id).First(&user)
	return user, nil
}

// done
// get email
func (r *Repo) GetUserEmail(email string) (model.User, error) {
	// GetUserEmail là hàm lấy thông tin user
	var user model.User
	r.db.Where("email = ?", email).First(&user)
	return user, nil
}

// done
// get username
func (r *Repo) GetUserByUsername(username string) (model.User, error) {
	// GetUserByUsername là hàm lấy thông tin user
	var user model.User
	r.db.Where("username = ?", username).First(&user)
	return user, nil
}

// get role
func (r *Repo) GetUserRole(role string) (model.User, error) {
	// GetUserRole là hàm lấy thông tin user
	var user model.User
	r.db.Where("role = ?", role).First(&user)
	return user, nil
}

// get address
func (r *Repo) GetUserAddress(address string) (model.User, error) {
	// GetUserAddress là hàm lấy thông tin user
	var user model.User
	r.db.Where("address = ?", address).First(&user)
	return user, nil
}

// get create user
func (r *Repo) GetUserCreateUser(create_user string) (model.User, error) {
	// GetUserCreateUser là hàm lấy thông tin user
	var user model.User
	r.db.Where("create_user = ?", create_user).First(&user)
	return user, nil
}

// get update user
func (r *Repo) GetUserUpdateUser(update_user string) (model.User, error) {
	// GetUserUpdateUser là hàm lấy thông tin user
	var user model.User
	r.db.Where("update_user = ?", update_user).First(&user)
	return user, nil
}

// get create time
func (r *Repo) GetUserCreateTime(create_time string) (model.User, error) {
	// GetUserCreateTime là hàm lấy thông tin user
	var user model.User
	r.db.Where("create_time = ?", create_time).First(&user)
	return user, nil
}

// get delete user
func (r *Repo) GetUserDeleteUser(delete_user int) (model.User, error) {
	// GetUserDeleteUser là hàm lấy thông tin user
	var user model.User
	r.db.Where("delete_user = ?", delete_user).First(&user)
	return user, nil
}

// get phone
func (r *Repo) GetUserPhone(phone string) (model.User, error) {
	// GetUserPhone là hàm lấy thông tin user
	var user model.User
	r.db.Where("phone = ?", phone).First(&user)
	return user, nil
}
