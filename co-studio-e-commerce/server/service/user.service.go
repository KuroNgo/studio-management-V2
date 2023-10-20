package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
	"co-studio-e-commerce/util"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	repo repo.IRepo
	user model.User
}

type IUser interface {
	GetAllUser() ([]model.User, error)
	FindUserByID(uuid string) (*model.User, error)
	FindUserByRole(role string) (*[]model.User, error)
	GetUserByID(uuid uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByRole(role string) (*[]model.User, error)
	GetUserByAddress(address string) (*[]model.User, error)
	GetUserCreateUser(created_at time.Time) (*[]model.User, error)
	LoginUserByEmail(UserRequest model.SignInInput) (userResponse *model.User, err error)
	LoginUserByUsername(UserRequest model.UserRequest) (UserResponse *model.User, err error)
	UpdateUser(currentUser model.User) (*model.User, error)
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

func (u *User) FindUserByID(uuid string) (*model.User, error) {
	user, err := u.repo.FindUserByID(uuid)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (u *User) FindUserByRole(role string) (*[]model.User, error) {
	user, err := u.repo.GetUserRole(role)
	if err != nil {
		return &[]model.User{}, err
	}
	return user, nil
}

func (u *User) GetUserByID(uuid uuid.UUID) (*model.User, error) {

	user, err := u.repo.GetUserID(uuid)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (u *User) GetUserByEmail(email string) (*model.User, error) {
	user, err := u.repo.GetUserEmail(email)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	user, err := u.repo.GetUserByUsername(username)
	if err != nil {
		return &model.User{}, nil
	}
	return user, nil
}

func (u *User) GetUserByRole(role string) (*[]model.User, error) {
	user, err := u.repo.GetUserRole(role)
	if err != nil {
		return &[]model.User{}, nil
	}
	return user, nil
}

func (u *User) GetUserByAddress(address string) (*[]model.User, error) {
	user, err := u.repo.GetUserAddress(address)
	if err != nil {
		return &[]model.User{}, nil
	}
	return user, nil
}

func (u *User) GetUserCreateUser(created_at time.Time) (*[]model.User, error) {
	user, err := u.repo.GetUserCreatedAt(created_at)
	if err != nil {
		return &[]model.User{}, nil
	}
	return user, nil
}

// Đăng nhập theo email và password
func (u *User) LoginUserByEmail(UserRequest model.SignInInput) (userResponse *model.User, err error) {
	user, err := u.repo.GetUserEmail(UserRequest.Email)
	if err != nil {
		return &model.User{}, err
	}

	// Kiểm tra xem mật khẩu đã nhập có đúng với mật khẩu đã hash trong cơ sở dữ liệu không
	if err := util.VerifyPassword(user.Password, UserRequest.Password); err != nil {
		return &model.User{}, errors.New("Tài khoản hoặc mật khẩu không đúng !")
	}

	return user, nil
}

// Đăng nhập theo username
func (u *User) LoginUserByUsername(UserRequest model.UserRequest) (UserResponse *model.User, err error) {
	user, err := u.repo.GetUserByUsername(UserRequest.Username)
	if err != nil {
		return &model.User{}, err
	}

	// Kiểm tra xem mật khẩu đã nhập có đúng với mật khẩu đã hash trong cơ sở dữ liệu không
	if err := util.VerifyPassword(user.Password, UserRequest.Password); err != nil {
		return &model.User{}, errors.New("Tài khoản hoặc mật khẩu không đúng !")
	}
	return user, nil
}

func (u *User) UpdateUser(currentUser model.User) (*model.User, error) {
	// UpdateUser là hàm cập nhật thông tin user
	user, err := u.repo.UpdateUser(&currentUser)
	if err != nil {
		return &model.User{}, err
	}
	// Trả về thông tin người dùng sau khi cập nhật
	return user, nil
}

// Đăng ký tài khoản
func (u *User) RegisterUser(userRegister model.User) (userResponse model.User, err error) {
	userRequest := model.User{
		FullName:   userRegister.FullName,
		Username:   userRegister.Username,
		Email:      strings.ToLower(userRegister.Email),
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
