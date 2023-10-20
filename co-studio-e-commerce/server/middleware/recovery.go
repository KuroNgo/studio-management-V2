package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Mục đích của việc tạo hàm recovery là để tránh việc xảy ra lỗi panic làm cho
// server bị đứt đoạn
// Mặc định Gin đã hỗ trợ cho mình middleware recovery nhưng chỉ catch với lỗi 500
// Đối với lỗi khác không phải 500 thì mình phải bắt theo cách khác
// Nó giống như việc sử dụng try catch như các ngôn ngữ khác

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "there are some internal error",
					"message": err,
				})
			}
		}()
		c.Next()
	}
}
