package controller

import (
	"github.com/gin-gonic/gin"
)

func HttpHandlerPing(c *gin.Context) {
	c.JSON(Success())
}
