package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/errors"
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
