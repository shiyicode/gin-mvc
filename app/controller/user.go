package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/chuxinplan/gin-mvc/common/auth"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/gin-gonic/gin"
)

// If `GET`, only `Form` binding engine (`query`) used.
// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).

func HttpHandlerLogin(c *gin.Context) {
	param := new(service.LoginParam)
	if err := c.ShouldBind(param); err != nil {
		panic(errors.Warp(errors.ErrValidation, err.Error()))
	}

	userService := service.NewUserService(getUsername(c), getRequestId(c))
	ret := userService.Login(param)

	token, err := auth.EncodeToken("test", 1)
	if err != nil {
		panic(errors.Warp(errors.ErrGetTokenFail, err.Error()))
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(Success(ret))
}

func HttpHandlerRegister(c *gin.Context) {
	c.JSON(Success())
}
