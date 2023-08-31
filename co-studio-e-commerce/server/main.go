package main

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/route"
)

func main() {
	conf.SetEnv()
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
