package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSPublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Next()
	}
}

func CORSForDev() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Methods", "GET, POST")
		c.Header("Access-Control-Max-Age", "86400") //24h

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Access-Control-Request-Method")
		c.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Max-Age", "86400") // 24 giờ

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func OptionMessage(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	c.Header("Access-Control-Max-Age", "86400") // 24 giờ

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusFound, 204)
	}

	c.Next()
}
