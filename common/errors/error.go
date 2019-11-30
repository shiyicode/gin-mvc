package errors

import (
	"fmt"
	"net/http"
)

// 使用 错误码 和 error 创建新的 错误
func Warp(errno *Errno, msg ...string) *Err {
	innerMsg := ""
	if len(msg) > 0 {
		innerMsg = msg[0]
	}
	return &Err{
		errno,
		innerMsg,
	}
}

type Err struct {
	*Errno          // 自定义错误码类型
	innerMsg string // 保存内部错误信息
}

func (err *Err) HTTPCode() int {
	if err.Errno.httpCode != 0 {
		return err.Errno.httpCode
	}
	return http.StatusInternalServerError
}

// 返回错误码
func (err *Err) Code() int {
	return err.Errno.code
}

func (err *Err) Error() string {
	return fmt.Sprintf("errno: %d | message: %s | error: %s", err.Errno.code, err.Errno.message, err.innerMsg)
}

func (err *Err) Message() string {
	msg := err.Errno.message

	switch err.Errno.code {
	case ErrBadRequest.code:
		msg += "; " + err.innerMsg
	}

	return msg
}
