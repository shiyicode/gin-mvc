package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpHandlerPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
