package main

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/route"
	"fmt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   Cỏ Studio API
// @version  1.0
// @description Đây là API của Cỏ Studio,
// @description Việc sử dụng API này phải có sự đồng ý của Mr. Phong
// @description Swagger chỉ phục vụ cho việc hiển thị các API được phép sử dụng
// @description Nếu bạn muốn sử dụng phải thực hiện trên postman
// @description Để có file test postman, bạn cần liên hệ người quản lý API

// @contact.name Ngô Hoài Phong
// @contact.url http://www.swagger.io/support
// @contact.email hoaiphong01012002@gmai.com

// @host localhost:8000/
func init() {
	conf.SetEnv()
	app := route.NewService()
	fmt.Println("Đường dẫn liên kết đến Swagger: " + "http://localhost:8000")
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func main() {
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.DeepLinking(true),
		ginSwagger.PersistAuthorization(true),
	)
}

//docker run -d -e POSTGRES_HOST=localhost -e POSTGRES_DB=kurodev -e POSTGRES_PASSWORD=01012002Phong.
//-e POSTGRES_PORT=8000 -e POSTGRES_USER=kurodev -p 5432:5432 --name co-studio-api postgres:latest
