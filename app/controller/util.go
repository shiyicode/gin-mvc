package controller

import (
	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/gin-gonic/gin"
)

func HttpHandlerPing(c *gin.Context) {
	userService := service.NewUserService(getUsername(c), getRequestId(c))

	userService.Logger.Infof("get pong")
	userService.Logger.Warningf("get pong")
	c.JSON(200, "ok")

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
