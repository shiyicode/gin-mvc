package middleware

import (
	"fmt"
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	confInfo := config.Get()
	if confInfo.Log.Enable {
		logPath := confInfo.Log.Path
		maxAge := time.Duration(confInfo.Log.MaxAge)
		rotationTime := time.Duration(confInfo.Log.RotateTime)
		writer := logger.GetLogWriter(logPath, "access.log", maxAge, rotationTime)

		loggerConf := gin.LoggerConfig{
			Output: writer,
			//SkipPaths: []string{"/test"},
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC1123),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.Request.UserAgent(),
					params.ErrorMessage,
				)
			},
		}
		return gin.LoggerWithConfig(loggerConf)
	}
	return gin.Logger()
}
