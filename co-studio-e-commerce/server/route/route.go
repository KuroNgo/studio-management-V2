package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/handler"
	"co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
)

type Service struct {
	*conf.App
}

type IRoute interface {
	NewService() *Service
}

func NewService() *Service {
	s := &Service{
		App: conf.NewApp(),
	}

	db := s.GetDB()
	repository := repo.NewRepo(db)

	// user
	// userService := service.NewUser(repository)
	// user := handler.NewUser(userService)

	// admin

	// activity log
	activityLogService := service.NewActivityLog(repository)
	activityLog := handler.NewActivityLog(activityLogService)

	// migration
	// migrate := handler.NewMigration(db)

	route := s.Router
	v1 := route.Group("/api/v1")

	// user
	v1.GET("/users", activityLog.GetActivityLogs)
	// route.POST("/migration", migra)
	return s
}
