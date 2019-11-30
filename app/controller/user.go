package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/app/service"
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

	userService := service.NewUserService(GetRequestId(c), GetDBSession(c))
	ret := userService.Login(param)

	cookie := &http.Cookie{
		Name:     "token",
		Value:    ret,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(Success(ret))
}

func HttpHandlerRegister(c *gin.Context) {
	param := new(service.RegisterParam)
	if err := c.ShouldBind(param); err != nil {
		panic(errors.Warp(errors.ErrValidation, err.Error()))
	}

	userService := service.NewUserService(GetRequestId(c), GetDBSession(c))
	userService.Register(param)

	c.JSON(Success())
}
