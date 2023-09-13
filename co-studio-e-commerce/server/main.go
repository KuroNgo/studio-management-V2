package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/route"
)

var (
	server *gin.Engine
)

func init() {
	conf.SetEnv()
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func main() {
	conf.DbDefault.Exec("CREATE EXTENSION IF NOT EXISTS \\\"uuid-ossp\\\"")
	conf.DbDefault.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Categories{},
		&model.Order{},
		&model.OrderDetail{},
		&model.Admin{},
		&model.ActivityLog{},
		&model.Login_Session{},
		&model.PaymentInformation{},
		&model.Post{},
		&model.Server{},
		&model.ShoppingCart{},
		&model.Reviews{},
	)
	fmt.Print("Migrate database success")

}
