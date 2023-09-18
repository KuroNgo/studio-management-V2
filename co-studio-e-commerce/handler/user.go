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

func NewUser(service service.IUser) *User {
	return &User{service: service}
}

func (u *User) GetAllUser(ctx *gin.Context) {
	users, err := u.service.GetAllUser(u.user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": users}})
}

// GetUser godoc
// @Summary Thực hiện tìm kiếm thông tin người dùng theo ID
// @Description Nhận thông tin chi tiết của người dùng hiện được xác thực
// @Accept json
// @Produce json
// @Router /api/v1/get/user [get]
func (u *User) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.User)

	userResponse := &model.UserResponse{
		ID:        currentUser.UserID,
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

// Login godoc
// @Summary Đăng nhập người dùng
// @Description Thực hiện chức năng đăng nhập bằng username
// @Accept json
// @Produce json
// @Router /api/v1/login/username [post]
func (u *User) LoginWithUserName(ctx *gin.Context) {
	// Lấy thông tin từ request
	userRequest := model.UserRequest{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	data, err := u.service.LoginUserByUsername(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Trả về thông báo login thành công
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful", "user": data})
}

// Login godoc
// @Summary Đăng nhập người dùng
// @Description Thực hiện chức năng đăng nhập bằng email
// @Accept json
// @Produce json
// @Router /api/v1/login/email [post]
func (u *User) LoginWithEmail(ctx *gin.Context) {
	//  Lấy thông tin từ request
	userRequest := model.SignInInput{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	data, err := u.service.LoginUserByEmail(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	// Trả về thông báo login thành công
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful", "user": data})
}

// Đăng ký người dùng godoc
// @Summary Đăng ký người dùng trên trang web
// @Description Hiển thị form đăng ký cho người dùng điền thông tin
// @Accept json
// @Produce json
// @Router /api/v1/register [post]
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

	if !util.PhoneValid(user.Phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Số điện thoại không đúng chuẩn"})
		return
	}

	// Bên phía client sẽ phải so sánh password thêm một lần nữa đã đúng chưa
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
