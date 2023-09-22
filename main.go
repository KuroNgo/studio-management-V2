package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/route"
)

var (
	server *gin.Engine
)

// @title   Cỏ Studio API
// @version  1.0
// @description Đây là API của Cỏ Studio,
// @description Việc sử dụng API này phải có sự đồng ý của Mr. Phong
// @description Swagger chỉ phục vụ cho việc hiển thị các API được phép sử dụng
// @description Nếu bạn muốn sử dụng phải thực hiện trên postman
// @description Để có file test postman, bạn cần liên hệ người quản lý API
// @termsOfService http://localhost:8000

// @contact.name Ngô Hoài Phong
// @contact.url http://www.swagger.io/support
// @contact.email hoaiphong01012002@gmai.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

//  @query.collection.format multi

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @host localhost:8000
// @query.collection.format multi
func init() {
	conf.SetEnv()
	app := route.NewService()

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
