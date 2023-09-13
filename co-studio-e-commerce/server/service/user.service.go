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
	LoginUserByUsername(UserRequest model.UserRequest) (userResponse model.User, err error)
	LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error)
	RegisterUser(request model.User) (userResponse model.User, err error)
}

func NewUser(repo repo.IRepo) *User {
	return &User{repo: repo}
}

// Đăng nhập theo email và password
func (u *User) LoginUserByEmail(UserRequest model.SignInInput) (userResponse model.User, err error) {
	user, err := u.repo.GetUserEmail(UserRequest.Password)
	if err != nil {
		return model.User{}, err
	}

	// Kiểm tra xem mật khẩu  có đúng không
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

// Đăng ký tài khoản
func (u *User) RegisterUser(userRegister model.User) (userResponse model.User, err error) {
	userRequest := model.User{
		Name:      userRegister.Name,
		Email:     userRegister.Email,
		Password:  userRegister.Password,
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
