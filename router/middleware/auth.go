package middleware

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/auth"
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
			c.JSON(http.StatusForbidden, "权限检验失败，请重新登录!")
			c.Abort()
			return
		}
		c.Set("userId", payLoad.UserId)
		c.Next()
	}
}
