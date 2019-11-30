package middleware

import (
	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/auth"
	"github.com/chuxinplan/gin-mvc/common/logger"
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		payLoad, err := auth.DecodeToken(token)
		if err == nil {
			c.Set("userId", payLoad.UserId)
		}
		c.Next()
	}
}

func MustGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		payLoad, err := auth.DecodeToken(token)
		if err != nil {
			reqLog := logger.GetRequestLogger(controller.GetRequestId(c))
			reqLog.Error(err.Error())
			c.JSON(controller.Failure(err))
			c.Abort()
			return
		}
		c.Set("userId", payLoad.UserId)
		c.Set("userName", payLoad.Username)
		c.Next()
	}
}
