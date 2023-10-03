package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/docs"
	"co-studio-e-commerce/handler"
	"co-studio-e-commerce/middleware"
	repo "co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
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

	//categoryService := service.NewCategoryService(repository)
	//category := handler.NewCategory(categoryService)

	route := s.Router
	docs.SwaggerInfo.BasePath = "api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//route automatically
	//Thực hiện tự động chuyển hướng khi chạy chương trình
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	v1User := route.Group("/user/v1")

	// auth
	v1User.POST("/login/username", user.LoginWithUserName)
	v1User.POST("/login/email", func(ctx *gin.Context) {
		user.LoginWithEmail(ctx, &conf.Config{})
	})
	v1User.GET("/get-all-user", user.GetAllUser)
	v1User.POST("/register", user.Register)
	v1User.GET("/get/user", user.GetMe)
	v1User.GET("/logout", middleware.DeserializeUser(), user.LogoutUser)
	//v1User.PUT("/update/user", user.UpdateUser)

	// category

	return &s
}
