package service

import (
	"errors"

	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
)

type User struct {
	repo repo.IRepo
}

type IUser interface {
	GetAllUser(user model.User) ([]model.User, error)
	GetUser(user *model.User) error
	LoginUserByUsername(UserRequest model.UserRequest) (userResponse model.User, err error)
	LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error)
	RegisterUser(request model.User) (userResponse model.User, err error)
}

func NewUser(repo repo.IRepo) *User {
	return &User{repo: repo}
}

func (u *User) GetUser(user *model.User) error {
	// GetUser là hàm lấy thông tin user
	if err := u.repo.GetUser(user); err != nil {
		return err
	}
	return nil
}

func (u *User) GetAllUser(user model.User) ([]model.User, error) {
	// GetAllUser là hàm lấy thông tin tất cả user
	users, err := u.repo.GetAllUser(user)
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

// Sai logic
// Đăng nhập theo email và password
func (u *User) LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error) {
	user, err := u.repo.GetUserEmail(UserRequest.Email)
	if err != nil {
		return model.User{}, err
	}

	// Kiểm tra xem mật khẩu  có đúng so với mật khẩu trong database không
	if user.Password != UserRequest.Password {
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

	// Kiểm tra xem mật khẩu  có đúng không
	if user.Password != UserRequest.Password {
		return model.User{}, errors.New("Tài khoản hoặc mật khẩu không đúng !")
	}
	return user, nil
}

// Sai logic
// Đăng ký tài khoản
func (u *User) RegisterUser(userRegister model.User) (userResponse model.User, err error) {
	userRequest := model.User{
		Name:      userRegister.Name,
		Email:     userRegister.Email,
		Password:  userRegister.Password,
		Phone:     userRegister.Phone,
		Role:      userRegister.Role,
		Provider:  userRegister.Provider,
		Photo:     userRegister.Photo,
		Verified:  userRegister.Verified,
		CreatedAt: userRegister.CreatedAt,
		UpdatedAt: userRegister.UpdatedAt,
		Enable:    userRegister.Enable,
	}
	user, err := u.repo.CreateUser(userRequest)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
