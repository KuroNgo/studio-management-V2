package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/docs"
	"co-studio-e-commerce/handler"
	"co-studio-e-commerce/middleware"
	repo "co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	//categoryService := service.NewCategoryService(repository)
	//category := handler.NewCategory(categoryService)

	route := s.Router
	route.MaxMultipartMemory = 25 << 20 // 8 MiB

	docs.SwaggerInfo.BasePath = "api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//route automatically
	//Thực hiện tự động chuyển hướng khi chạy chương trình
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})
	v1User := route.Group("/user/v1")
	{

		v1User.Use(middleware.CORSForDev())

		// auth
		v1User.POST("/login/username", user.LoginWithUserName)
		v1User.POST("/login/email", user.LoginWithEmail)
		v1User.GET("/get-all-user", user.GetAllUser)
		v1User.POST("/register", user.Register)
		v1User.GET("/get/user", middleware.DeserializeUser(), middleware.ProtectedCurrentUser(), user.GetMe)
		v1User.GET("/logout", middleware.DeserializeUser(), user.LogoutUser)
		v1User.GET("/refresh", user.RefreshAccessToken)
		//v1User.PUT("/update/user", user.UpdateUser)

		// image
		v1User.GET("image/get", util.GetUploadedFile)
		v1User.POST("/image/upload", util.UploadAFile)
		v1User.PUT("/images/uploads", util.UpdateFile)
		v1User.DELETE("/image/delete", util.DeleteFile)
		v1User.POST("/image/", util.UploadMultiFile)
	}

	// category

	return &s
}
