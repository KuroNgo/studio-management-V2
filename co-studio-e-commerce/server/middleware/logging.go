package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var db *gorm.DB
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = ctx.Writer.Size()

		if ctx.Errors != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Have err",
			})
			ctx.Next()
		}

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var logEvent *zerolog.Event
		if ctx.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)

		//newLog := model.ActivityLog{
		//	ClientIP:     param.ClientIP,
		//	Method:       param.Method,
		//	StatusCode:   param.StatusCode,
		//	BodySize:     param.BodySize,
		//	Path:         param.Path,
		//	Latency:      param.Latency.String(),
		//	ActivityTime: param.TimeStamp,
		//}
		//
		//db.Create(&newLog)
	}
}
