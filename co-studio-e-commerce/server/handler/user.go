package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
)

type User struct {
	service service.IUser
	user    model.User
}

type IUser interface {
}

func NewUser(service service.IUser) *User {
	return &User{service: service}
}

// GetMe là hàm lấy thông tin user hiện tại
func (u *User) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.User)

	userResponse := &model.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// Đăng nhập
func (u *User) Login(ctx *gin.Context) {
	// Lấy thông tin từ request
	userRequest := model.UserRequest{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	return
}

// Đăng ký
func (u *User) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if !util.EmailValid(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Email không hợp lệ !"})
		return
	}

	if !util.PhoneValid(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Mật khẩu phải có ít nhất 8 ký tự, bao gồm 1 chữ hoa, 1 chữ thường và 1 số !"})
		return
	}

	password := user.Password
	if !util.PasswordStrong(password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Mật khẩu phải có ít nhất 8 ký tự, có thể bao gồm 1 chữ hoa, hoặc 1 chữ thường hoặc 1 số hoặc bao gồm cả 3 !"})
		return
	}

	// Băm mật khẩu
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	user.Password = hashedPassword

	// thực hiện đăng ký người dùng
	userResponse, err := u.service.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Trả về phản hồi thành công
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
