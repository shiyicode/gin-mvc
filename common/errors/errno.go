package errors

import (
	"net/http"
)

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

type Errno struct {
	Code     int    // 错误码
	Message  string // 展示给用户看的
	HTTPCode int    // HTTP状态码
}

var (
	OK = &Errno{Code: 0, Message: "OK", HTTPCode: http.StatusOK}

	// 系统错误, 前缀为 100
	ErrInternalServer = &Errno{Code: 10001, Message: "内部服务器错误", HTTPCode: http.StatusInternalServerError}
	ErrParamConvert   = &Errno{Code: 10002, Message: "参数转换时发生错误", HTTPCode: http.StatusInternalServerError}

	// 数据库错误, 前缀为 201
	ErrDatabase = &Errno{Code: 20101, Message: "数据库错误", HTTPCode: http.StatusInternalServerError}

	// 认证错误, 前缀是 202
	ErrBind         = &Errno{Code: 20201, Message: "请求参数错误", HTTPCode: http.StatusBadRequest}
	ErrValidation   = &Errno{Code: 20202, Message: "验证失败", HTTPCode: http.StatusForbidden}
	ErrGetTokenFail = &Errno{Code: 20203, Message: "获取 token 失败", HTTPCode: http.StatusForbidden}

	// 用户错误, 前缀为 203
	ErrUserNotFound      = &Errno{Code: 20301, Message: "用户不存在", HTTPCode: http.StatusBadRequest}
	ErrPasswordIncorrect = &Errno{Code: 20302, Message: "密码错误", HTTPCode: http.StatusUnauthorized}
	ErrUserNotLogin      = &Errno{Code: 20303, Message: "用户未登陆", HTTPCode: http.StatusUnauthorized}
)
