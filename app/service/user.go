package service

import (
	"regexp"

	"github.com/chuxinplan/gin-mvc/common/auth"

	"github.com/chuxinplan/gin-mvc/app/model"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/go-xorm/xorm"
)

type UserService struct {
	baseService
}

type RegisterParam struct {
	Email    string `form:"email" binding:"required,email,max=200"`
	Username string `form:"username" binding:"required,max=200"`
	Password string `form:"password" binding:"required,max=200"`
}

type LoginParam struct {
	Type     string `form:"type" binding:"required,oneof=email username"`
	Account  string `form:"account" binding:"required,max=200"`
	Password string `form:"password" binding:"required,max=200"`
}

func NewUserService(requestId string, db *xorm.Session) UserService {
	return UserService{
		baseService: newBaseService(requestId, db),
	}
}

func (userService UserService) Login(param *LoginParam) string {
	var userInfo *model.User
	var err error
	if param.Type == "email" {
		userInfo, err = model.UserGetByEmail(userService.DB, param.Account)
		if err != nil {
			panic(errors.Warp(errors.ErrDatabase, err.Error()))
		}
	} else if param.Type == "username" {
		userInfo, err = model.UserGetByUsername(userService.DB, param.Account)
		if err != nil {
			panic(errors.Warp(errors.ErrDatabase, err.Error()))
		}
	}

	if userInfo == nil {
		panic(errors.Warp(errors.ErrUserNotFound))
	}

	if userInfo.Password != param.Password {
		panic(errors.Warp(errors.ErrPasswordIncorrect))
	}

	token, err := auth.EncodeToken(userInfo.Username, userInfo.Id)
	if err != nil {
		panic(errors.Warp(errors.ErrGetTokenFail, err.Error()))
	}

	return token
}

func (userService UserService) Register(param *RegisterParam) {
	// 用户名由字母数据下划线英文句号组成，长度要求4-16之间
	usernameReg, err := regexp.Compile(`^[a-zA-Z0-9_\.]{4,16}$`)
	if err != nil {
		panic(errors.Warp(errors.ErrInternalServer, err.Error()))
	}
	if usernameReg.MatchString(param.Username) == false {
		panic(errors.Warp(errors.ErrUsernameValidation))
	}

	// 密码匹配6-16位英文数据大部分英文标点
	passwordReg, err := regexp.Compile(`^([A-Za-z0-9\-=\[\];,\./~!@#\$%^\*\(\)_\+}{:\?]){6,16}$`)
	if err != nil {
		panic(errors.Warp(errors.ErrInternalServer, err.Error()))
	}
	if passwordReg.MatchString(param.Password) == false {
		panic(errors.Warp(errors.ErrPasswordValidation))
	}
	// 密码至少包含一个大写英文
	passwordReg, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		panic(errors.Warp(errors.ErrInternalServer, err.Error()))
	}
	if passwordReg.MatchString(param.Password) == false {
		panic(errors.Warp(errors.ErrPasswordValidation))
	}
	// 密码至少包含一个小写英文
	passwordReg, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		panic(errors.Warp(errors.ErrInternalServer, err.Error()))
	}
	if passwordReg.MatchString(param.Password) == false {
		panic(errors.Warp(errors.ErrPasswordValidation))
	}
	// 密码至少包含一个小写英文
	passwordReg, err = regexp.Compile(`[\-=\[\];,\./~!@#\$%^\*\(\)_\+}{:\?]+`)
	if err != nil {
		panic(errors.Warp(errors.ErrInternalServer, err.Error()))
	}
	if passwordReg.MatchString(param.Password) == false {
		panic(errors.Warp(errors.ErrPasswordValidation))
	}

	// 判断Username唯一性
	userInfo, err := model.UserGetByUsername(userService.DB, param.Username)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	if userInfo != nil {
		panic(errors.Warp(errors.ErrUserRegisterAgain))
	}

	// 判断email唯一性
	userInfo, err = model.UserGetByEmail(userService.DB, param.Email)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	} else if userInfo != nil {
		panic(errors.Warp(errors.ErrUserRegisterAgain))
	}

	createUserInfo := &model.User{
		Email:    param.Email,
		Username: param.Username,
		Password: param.Password,
	}
	_, err = model.UserCreate(userService.DB, createUserInfo)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
}
