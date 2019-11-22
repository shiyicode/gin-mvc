package middleware

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/auth"
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func MustGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		userId, err := checkLogin(token)
		if err != nil {
			c.JSON(http.StatusForbidden, "权限检验失败，请重新登录!")
			c.Abort()
		}
		c.Set("userId", userId)
		c.Next()
	}
}

//input token，get decode userId
func checkLogin(token string) (int64, error) {
	data, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	token = string(data)
	if token == "" {
		return 0, errors.New("token串为空!")
	}

	isLogin,payLoad := auth.DecodeToken(token)
	if isLogin{
		return payLoad.UserId,nil
	}

	return 0, errors.New("用户会话有误或已失效！")
}
