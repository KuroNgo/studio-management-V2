package handler

import (
	"co-studio-e-commerce/conf"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	users, err := u.service.GetAllUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"users": users}})
}

func (u *User) GetUserByID(ctx *gin.Context) {

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
		ID:        currentUser.ID,
		Name:      currentUser.Username,
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

	cfg, _ := conf.LoadConfig(".")

	// Generate token
	access_token, err := util.CreateToken(cfg.AccessTokenExpiresIn, data.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := util.CreateToken(cfg.RefreshTokenExpiresIn, data.ID, cfg.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, cfg.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

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

	cfg, _ := conf.LoadConfig(".")

	// Generate token
	access_token, err := util.CreateToken(cfg.AccessTokenExpiresIn, data.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := util.CreateToken(cfg.RefreshTokenExpiresIn, data.ID, cfg.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, cfg.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

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

	user.Password = util.Santize(user.Password)
	user.Username = util.Santize(user.Username)
	user.Email = util.Santize(user.Email)
	user.Email = strings.ToLower(user.Email)

	// Bên phía client sẽ phải so sánh password thêm một lần nữa đã đúng chưa
	if !util.PasswordStrong(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Mật khẩu phải có ít nhất 8 ký tự, bao gồm chữ hoa, chữ thường và số !"})
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

func (u *User) RefreshAccessToken(ctx *gin.Context) {
	message := "could not fresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	cfg, _ := conf.LoadConfig(".")

	sub, err := util.ValidateToken(cookie, cfg.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result, err := u.service.FindUserByID(fmt.Sprint(sub))
	resultString, err := json.Marshal(result)
	userResponse := &model.UserResponse{
		ID:        result.ID,
		Name:      result.Username,
		Email:     result.Email,
		Photo:     result.Photo,
		Role:      result.Role,
		Provider:  result.Provider,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": string(resultString) + "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := util.CreateToken(cfg.AccessTokenExpiresIn, result.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"access_token": access_token,
		"user":         userResponse,
	})
}

func (u *User) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}