package middleware

import (
	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/auth"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		isLogin,payLoad := auth.DecodeToken(token)
		if isLogin{
			c.Set("userId", payLoad.UserId)
		}
		c.Next()
	}
}

func MustGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		isLogin,payLoad := auth.DecodeToken(token)
		if !isLogin {
			resErr := errors.New(errors.ErrValidation, "")
			httpCode,respData := controller.Failure(resErr)
			c.JSON(httpCode, respData)
			c.Abort()
			return
		}
		c.Set("userId", payLoad.UserId)
		c.Next()
	}
}
