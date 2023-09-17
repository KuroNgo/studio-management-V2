package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/docs"
	"co-studio-e-commerce/handler"
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

	route := s.Router
	docs.SwaggerInfo.BasePath = "api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//route automatically
	//Thực hiện tự động chuyển hướng khi chạy chương trình
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	v1 := route.Group("/api/v1")

	// auth
	v1.POST("/login/username", user.LoginWithUserName)
	v1.POST("/login/email", user.LoginWithEmail)
	v1.POST("/register", user.Register)
	v1.GET("/get/user", user.GetMe)

	return &s
}
