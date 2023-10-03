package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
	"co-studio-e-commerce/util"
	"errors"
)

type User struct {
	repo repo.IRepo
	user model.User
}

type IUser interface {
	GetAllUser() ([]model.User, error)
	LoginUserByUsername(UserRequest model.UserRequest) (userResponse model.User, err error)
	LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error)
	UpdateUser(currentUser model.User, user model.User) (model.User, error)
	RegisterUser(userRegister model.User) (userResponse model.User, err error)
}

func NewUser(repo repo.IRepo) *User {
	return &User{repo: repo}
}

func (u *User) GetAllUser() ([]model.User, error) {
	// GetAllUser là hàm lấy thông tin tất cả user
	users, err := u.repo.GetAllUser()
	if err != nil {
		return nil, err // Trả về nil và lỗi nếu có lỗi
	}
	return users, nil // Trả về danh sách người dùng và không có lỗi nếu thành công
}

// Đăng nhập theo email và password
func (u *User) LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error) {
	user, err := u.repo.GetUserEmail(UserRequest.Email)
	if err != nil {
		return model.User{}, err
	}

	// Kiểm tra xem mật khẩu đã nhập có đúng với mật khẩu đã hash trong cơ sở dữ liệu không
	if err := util.VerifyPassword(user.Password, UserRequest.Password); err != nil {
		return model.User{}, errors.New("Tài khoản hoặc mật khẩu không đúng !")
	}

	return user, nil
}

// Đăng nhập theo username
func (u *User) LoginUserByUsername(UserRequest model.UserRequest) (UserResponse model.User, err error) {
	user, err := u.repo.GetUserByUsername(UserRequest.Username)
	if err != nil {
		return model.User{}, err
	}

	// Kiểm tra xem mật khẩu đã nhập có đúng với mật khẩu đã hash trong cơ sở dữ liệu không
	if err := util.VerifyPassword(user.Password, UserRequest.Password); err != nil {
		return model.User{}, errors.New("Tài khoản hoặc mật khẩu không đúng !")
	}
	return user, nil
}

func (u *User) UpdateUser(currentUser model.User, userRequest model.User) (model.User, error) {
	// UpdateUser là hàm cập nhật thông tin user
	user, err := u.repo.UpdateUser(currentUser, &userRequest)
	if err != nil {
		return model.User{}, err
	}
	// Trả về thông tin người dùng sau khi cập nhật
	return user, nil
}

// Đăng ký tài khoản
func (u *User) RegisterUser(userRegister model.User) (userResponse model.User, err error) {
	userRequest := model.User{
		FullName:   userRegister.FullName,
		Username:   userRegister.Username,
		Email:      userRegister.Email,
		Password:   userRegister.Password,
		Phone:      userRegister.Phone,
		Role:       userRegister.Role,
		Provider:   userRegister.Provider,
		AvatarUser: userRegister.AvatarUser,
		Photo:      userRegister.Photo,
		Enable:     userRegister.Enable,
	}
	user, err := u.repo.CreateUser(userRequest)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// hook
