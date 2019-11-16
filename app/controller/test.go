package controller

import (
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func HttpHandlerTest(c *gin.Context) {
	log.Infof("get pong\n")
	log.Warningf("get pong\n")
	c.JSON(200, gin.H{
		"message": "pong",
	})

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
