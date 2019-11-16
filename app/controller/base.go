package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/error"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// data为可选参数
func Success(c *gin.Context, data ...interface{}) {
	result := &Result{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
	if len(data) > 0 {
		result.Data = data[0]
	}
	c.JSON(http.StatusOK, result)
}

func Failure(c *gin.Context, err *error.Err) {
	//panic(&HttpResponse{1, msg[0], 0})
	result := &Result{
		Code:    err.Code(),
		Message: err.Error(),
		Data:    nil,
	}
	c.JSON(err.HTTPCode(), result)
}

func GetUserId(c *gin.Context) int64 {
	userId, exists := c.Get("userId")
	if exists == false {
		return 0
	}
	id, ok := userId.(int64)
	if !ok {
		return 0
	}
	return id
}
