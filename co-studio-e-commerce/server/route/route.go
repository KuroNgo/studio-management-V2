package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/docs"
	"co-studio-e-commerce/handler"
	"co-studio-e-commerce/middleware"
	repo "co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Service struct {
	*conf.App
}

func NewService() *Service {
	s := Service{
		conf.NewApp(),
	}

	db := s.GetDB()
	repository := repo.NewRepo(db)

	userService := service.NewUser(repository)
	user := handler.NewUser(userService)

	categoryService := service.NewCategoryService(repository)
	category := handler.NewCategory(categoryService)

	route := s.Router
	route.Use(middleware.CORSMiddleware())

	route.MaxMultipartMemory = 25 << 20 // 8 MiB

	docs.SwaggerInfo.BasePath = "api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//route automatically
	//Thực hiện tự động chuyển hướng khi chạy chương trình
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	// option method to fix preflight request
	route.OPTIONS("/*path", middleware.OptionMessage)

	v1User := route.Group("api/user/v1")
	{

		v1User.Use(
			middleware.CORSForDev(),
			middleware.RateLimiter(),
		)

		// auth

	}

	clientV1 := route.Group("api/client/v1")
	{
		clientV1.Use(middleware.CORSMiddleware())
		clientV1.Use(middleware.RateLimiter())

		// auth
		clientV1.POST("/login/username", user.LoginWithUserName)
		clientV1.POST("/login/email", user.LoginWithEmail)
		clientV1.POST("/register", user.Register)
		clientV1.GET("/get-user", user.GetMeV2)
		clientV1.GET("/refresh", user.RefreshAccessToken)
		clientV1.GET("/logout", middleware.DeserializeUser(), user.LogoutUser)
		clientV1.PUT("/update", user.UpdateUser)

		// category

		// image
		clientV1.GET("/image/get", util.GetUploadedFile)
		clientV1.POST("/image/upload", util.UploadAFile)
		clientV1.PUT("/images/uploads", util.UpdateFile)
		clientV1.DELETE("/image/delete", util.DeleteFile)
		clientV1.POST("/image/", util.UploadMultiFile)
	}

	// Đối với admin thì không cần phải rate limit
	// Nhưng đối với nhân viên thì khác
	// Cơ chế rate limit sẽ giới hạn đối với nhân viên
	// Nhưng số lần gửi request giới hạn sẽ được tăng lên so với người dùng user
	adminV1 := route.Group("api/admin/v1")
	{
		adminV1.Use(
			middleware.CORSForDev(),
			middleware.DeserializeUser(),
		)

		// user
		adminV1.GET("/get-all-user", user.GetAllUser)
		adminV1.GET("/get-role", user.GetUserByRole)

		// category
		adminV1.POST("/category/post", category.CreateCategory)
	}

	return &s
}
