package controller

import (
	"encoding/base64"
	"net/http"

	"github.com/chuxinplan/gin-mvc/app/model"
	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/chuxinplan/gin-mvc/router/middleware"
	"github.com/gin-gonic/gin"
)

// If `GET`, only `Form` binding engine (`query`) used.
// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).

func HttpHandlerLogin(c *gin.Context) {
	param := new(service.LoginParam)
	if err := c.ShouldBind(param); err != nil {
		panic(errors.New(errors.ErrValidation, err.Error()))
	}

	userService := service.NewUserService(getUsername(c), getRequestId(c))
	ret := userService.Login(param)

	userMess := &model.User{
		Id :1,
		Username:"test",
	}
	token := middleware.GetToken(userMess)
	cookie := &http.Cookie{
		Name:     "token",
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(success(ret))
}

func HttpHandlerRegister(c *gin.Context) {
	c.JSON(success())
}
