package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/handler"
	repo "co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
)

type Service struct {
	*conf.App
}

type IRoute interface {
	NewService() *Service
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
	v1 := route.Group("/api/v1")

	// auth
	v1.POST("/login", user.Login)
	v1.POST("/register", user.Register)

	return &s
}
