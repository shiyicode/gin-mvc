package middleware

import (
	"fmt"
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Logger() gin.HandlerFunc {
	confInfo := config.Get()
	if confInfo.Log.Enable {
		logPath := confInfo.Log.Path
		maxAge := time.Duration(confInfo.Log.MaxAge)
		rotationTime := time.Duration(confInfo.Log.RotateTime)
		writer := logger.GetLogWriter(logPath, "access.log", maxAge, rotationTime)

		return func(ctx *gin.Context) {
			requestId := ctx.Request.Header.Get("X-Request-Id")

			if requestId == "" {
				requestId = uuid.NewV4().String()
			}
			ctx.Set("X-Request-Id", requestId)
			ctx.Header("X-Request-Id", requestId)

			startTime := time.Now()
			ctx.Next()
			endTime := time.Now()
			latencyTime := endTime.Sub(startTime)

			// 日志格式
			fmt.Fprintf(writer, "%15s - %s %s \"%s %s %s %3d %13v \"%s\n",
				ctx.ClientIP(),
				startTime.Format("2006-01-02 15:04:05"),
				requestId,
				ctx.Request.Method,
				ctx.Request.RequestURI,
				ctx.Request.Proto,
				ctx.Writer.Status(),
				latencyTime,
				ctx.Request.UserAgent(),
			)
		}
	}
	return gin.Logger()
}
