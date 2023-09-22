package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
)

func RetiverRatelimit(rateLimter *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ipClient := c.ClientIP()
		limiterCtx, err := rateLimter.Get(c, ipClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ratelimiter faile",
			})
			c.Abort()
			return
		}

		if limiterCtx.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "too many request",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
