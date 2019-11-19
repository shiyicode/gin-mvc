package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// data为可选参数
func success(data ...interface{}) (int, *Result) {
	result := &Result{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
	if len(data) > 0 {
		result.Data = data[0]
	}
	return http.StatusOK, result
}

func failure(err *errors.Err) (int, *Result) {
	result := &Result{
		Code:    err.Code(),
		Message: err.Error(),
		Data:    nil,
	}
	return err.HTTPCode(), result
}

func getUsername(c *gin.Context) string {
	username, exists := c.Get("username")
	if exists == false {
		return ""
	}
	name, ok := username.(string)
	if !ok {
		return ""
	}
	return name
}

func getRequestId(c *gin.Context) string {
	requestId, exists := c.Get("requestId")
	if exists == false {
		return ""
	}
	reqId, ok := requestId.(string)
	if !ok {
		return ""
	}
	return reqId
}
