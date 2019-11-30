package controller

import (
	"net/http"

	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// data为可选参数
func Success(data ...interface{}) (int, *Result) {
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

func Failure(err *errors.Err) (int, *Result) {
	result := &Result{
		Code:    err.Code(),
		Message: err.Message(),
		Data:    nil,
	}

	return err.HTTPCode(), result
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

func GetUsername(c *gin.Context) string {
	username, exists := c.Get("userName")
	if exists == false {
		return ""
	}
	name, ok := username.(string)
	if !ok {
		return ""
	}
	return name
}

func GetRequestId(c *gin.Context) string {
	requestId, exists := c.Get("X-Request-Id")
	if exists == false {
		requestId = ""
	}
	reqId, ok := requestId.(string)
	if !ok {
		return ""
	}
	return reqId
}

func GetDBSession(c *gin.Context) *xorm.Session {
	DBSession, exists := c.Get("db")
	if exists == false {
		return nil
	}
	db, ok := DBSession.(*xorm.Session)
	if !ok {
		return nil
	}
	return db
}
