package repo

import (
	"co-studio-e-commerce/middleware"
	"co-studio-e-commerce/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// done
func (r *Repo) GetUserProfile(email string) (model.User, error) {
	// GetUser là hàm lấy thông tin user
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("Người dùng không tồn tại")
		}
	}
	return user, nil
}

// Done
func (r *Repo) GetAllUser() ([]model.User, error) {
	// GetAllUser là hàm lấy thông tin tất cả user
	var users []model.User

	// Thực hiện truy vấn hoặc lấy thông tin tất cả người dùng từ nguồn dữ liệu của bạn
	if err := r.db.Find(&users).Order("role DESC").Error; err != nil {
		return nil, err // Trả về nil và lỗi nếu có lỗi
	}

	return users, nil // Trả về danh sách người dùng và không có lỗi nếu thành công
}

// Done
func (r *Repo) CreateUser(user model.User) (model.User, error) {
	// CreateUser là hàm tạo mới user

	// Thực hiện tạo mới user trong cơ sở dữ liệu
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err // Trả về lỗi nếu có lỗi
	}

	return user, nil // Trả về thông tin người dùng và không có lỗi nếu thành công
}

// cập nhật user, nếu user chưa có sẽ thực hiện thêm
func (r *Repo) UpdateUserORInsert(user *model.User) (model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return model.User{}, err
	}
	// Trả về thông tin người dùng sau khi cập nhật
	return *user, nil
}

// cập nhật người dùng
func (r *Repo) UpdateUser(currentUser model.User, user *model.User) (model.User, error) {
	// UpdateUser là hàm cập nhật thông tin user
	if err := r.db.Model(&currentUser).Where("userID = ?", currentUser.ID.String()).Updates(user).Omit("ID", "email").Error; err != nil {
		return model.User{}, err
	}

	// Trả về thông tin user đã được cập nhật
	return *user, nil
}

func (r *Repo) DeactivateUser(userID uuid.UUID, currentUser model.User) error {
	// Kiểm tra xem người thực hiện thao tác này có quyền (admin) hay không
	if !middleware.CurrentUserIsAdmin(currentUser) {
		return errors.New("Bạn không có quyền xóa người dùng.")
	}

	// Đánh dấu người dùng có ID là userID là "bị vô hiệu hóa" trong cơ sở dữ liệu
	if err := r.db.Model(&model.User{}).Where("userID = ?", userID).Update("enable", 0).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUser(user *model.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

// TODO: thay đổi trạng thái user
// thay đổi trạng thái có 2 tác động: admin hoặc tự động
// admin: thay đổi trạng thái của user
// tự động: thay đổi trạng thái của user khi hết hạn ( hoặc user bị dính black list)
//func (r *Repo) ChangeUserStatus(user *model.User, changeByAdmin bool) error {
//	// ChangeUserStatus là hàm thay đổi trạng thái user
//	if changeByAdmin {
//		// Trạng thái được thay đổi bởi admin
//		if err := r.db.Model(user).Update("status", user.Enable).Error; err != nil {
//			return err
//		}
//	} else {
//		// Trạng thái tự động thay đổi (ví dụ: khi hết hạn hoặc bị đưa vào blacklist)
//		// Thêm xử lý tương ứng ở đây
//	}
//
//	return nil
//}

// get id user
// login
func (r *Repo) GetUserID(userID uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.db.Where("userID = ?", userID.String()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("User not found")
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *Repo) FindUserByID(uuid string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "userid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// done
// get email
func (r *Repo) GetUserEmail(email string) (model.User, error) {
	// GetUserEmail là hàm lấy thông tin user
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("Email not found!")
		}
	}
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
func (r *Repo) GetUserRole(role string) ([]model.User, error) {
	// GetUserRole là hàm lấy thông tin user
	var user []model.User
	r.db.Where("role = ?", role).Find(&user)
	return user, nil
}

// get address
func (r *Repo) GetUserAddress(address string) ([]model.User, error) {
	// GetUserAddress là hàm lấy thông tin user
	var user []model.User
	r.db.Where("address = ?", address).Find(&user)
	return user, nil
}

// get create user
func (r *Repo) GetUserCreateUser(create_user string) ([]model.User, error) {
	// GetUserCreateUser là hàm lấy thông tin user
	var user []model.User
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
	r.db.Where("created_at = ?", create_time).First(&user)
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
