package errors

import (
	"net/http"
)

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三四位表示服务模块代码
第五六位表示具体错误代码
*/

type Errno struct {
	code     int    // 错误码
	message  string // 展示给用户看的
	httpCode int    // HTTP状态码
}

var (
	OK = &Errno{code: 0, message: "OK", httpCode: http.StatusOK}

	// 系统错误
	ErrUnKnown        = &Errno{code: 100000, message: "未知错误", httpCode: http.StatusInternalServerError}
	ErrInternalServer = &Errno{code: 100001, message: "内部服务器错误", httpCode: http.StatusInternalServerError}
	ErrParamConvert   = &Errno{code: 100002, message: "参数转换时发生错误", httpCode: http.StatusInternalServerError}
	ErrDatabase       = &Errno{code: 100003, message: "数据库错误", httpCode: http.StatusInternalServerError}

	// 模块通用错误
	ErrValidation      = &Errno{code: 200001, message: "参数校验失败", httpCode: http.StatusForbidden}
	ErrBadRequest      = &Errno{code: 200002, message: "请求参数错误", httpCode: http.StatusBadRequest}
	ErrGetTokenFail    = &Errno{code: 200003, message: "获取 token 失败", httpCode: http.StatusForbidden}
	ErrTokenNotFound   = &Errno{code: 200004, message: "用户 token 不存在", httpCode: http.StatusUnauthorized}
	ErrTokenExpire     = &Errno{code: 200005, message: "用户 token 过期", httpCode: http.StatusForbidden}
	ErrTokenValidation = &Errno{code: 200005, message: "用户 token 无效", httpCode: http.StatusForbidden}

	// User模块错误
	ErrUserNotFound       = &Errno{code: 200104, message: "用户不存在", httpCode: http.StatusBadRequest}
	ErrPasswordIncorrect  = &Errno{code: 200105, message: "密码错误", httpCode: http.StatusBadRequest}
	ErrUserRegisterAgain  = &Errno{code: 200107, message: "重复注册", httpCode: http.StatusBadRequest}
	ErrUsernameValidation = &Errno{code: 200107, message: "用户名不合法", httpCode: http.StatusBadRequest}
	ErrPasswordValidation = &Errno{code: 200107, message: "密码不合法", httpCode: http.StatusBadRequest}

	// Group模块错误
)
