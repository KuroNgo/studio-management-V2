package route

import (
	"co-studio-e-commerce/conf"
)

type Service struct {
	*conf.App
}

type IRoute interface {
	NewService() *Service
}

// func NewService() *Service {
// 	s := &Service{
// 		App: conf.NewApp(),
// 	}
//
// 	db := s.GetDB()
// 	repository := repo.NewRepo(db)
//
// }
