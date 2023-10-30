package middleware

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := conf.LoadConfig(".")
		sub, err := util.ValidateToken(access_token, config.AccessTokenPublicKey)

		if err != nil {
			fmt.Println("The err is: ", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		//var user model.User
		//
		//result := conf.DbDefault.First(&user, "userid = ?", fmt.Sprint(sub))
		//if result.Error != nil {
		//	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		//	return
		//}

		ctx.Set("currentUser", sub)
		ctx.Next()
	}
}
