package middleware

import (
	"fmt"

	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestId, _ := c.Get("X-Request-Id")

				if err, ok := err.(*errors.Err); ok {
					httpCode, respData := controller.Failure(err)
					c.JSON(httpCode, respData)
					errInfo := fmt.Sprintf("Err - errno: %d, message: %s, error: %s", err.Errno.Code, err.Errno.Message, err.InnerMsg)

					logger.Warningf("[%v] ErrorInfo:%s", requestId, errInfo)
					return
				}
				resErr := errors.Warp(errors.ErrInternalServer, "")
				httpCode, respData := controller.Failure(resErr)
				c.JSON(httpCode, respData)
				logger.Warningf("[%v] Unknown Error:%v", requestId, err)
			}
		}()
		c.Next()
	}
}
