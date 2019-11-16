package error

import (
	"fmt"
	"net/http"
)

// 使用 错误码 和 error 创建新的 错误
func New(errno *Errno, err error) *Err {
	return &Err{
		errno,
		err.Error(),
	}
}

type Err struct {
	*Errno          // 自定义错误码类型
	InnerMsg string // 保存内部错误信息
}

func (err *Err) HTTPCode() int {
	if err.Errno.HTTPCode != 0 {
		return err.Errno.HTTPCode
	}
	return http.StatusInternalServerError
}

// 返回错误码
func (err *Err) Code() int {
	return err.Errno.Code
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.InnerMsg)
}
