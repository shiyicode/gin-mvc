package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	confInfo := config.Get()

	var writer io.Writer
	if confInfo.Log.Enable {
		logPath := confInfo.Log.Path
		maxAge := time.Duration(confInfo.Log.MaxAge)
		rotationTime := time.Duration(confInfo.Log.RotateTime)
		writer = logger.GetLogWriter(logPath, "access.log", maxAge, rotationTime)
	} else {
		writer = os.Stdout
	}

	var logFormatter = formatter

	loggerConfig := gin.LoggerConfig{
		Formatter: logFormatter,
		Output:    writer,
		SkipPaths: []string{},
	}
	return gin.LoggerWithConfig(loggerConfig)
}

func formatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	requestId, exists := param.Keys["X-Request-Id"]
	if exists == false {
		requestId = ""
	}
	requestId, ok := requestId.(string)
	if !ok {
		requestId = ""
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s | %s | %s %-7s %s %s\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		requestId,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
