package controller

import (
	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/chuxinplan/gin-mvc/common/errors"
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
	c.JSON(success(ret))

	// token, err := managers.AccountLogin(account.Email, account.Password)
	// if err != nil {
	// 	c.JSON(http.StatusOK, base.Fail(err.Error()))
	// 	return
	// }
	// cookie := &http.Cookie{
	// 	Name:     "token",
	// 	Value:    base64.StdEncoding.EncodeToString([]byte(token)),
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }
	//
	// http.SetCookie(c.Writer, cookie)
}

func HttpHandlerRegister(c *gin.Context) {
	c.JSON(success())
}
