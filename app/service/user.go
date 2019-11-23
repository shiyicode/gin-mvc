package service

import "github.com/go-xorm/xorm"

type UserService struct {
	baseService
}

type RegisterParam struct {
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginParam struct {
	Type     string `form:"type" binding:"required,oneof=email username"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func NewUserService(requestId string, db *xorm.Session) UserService {
	return UserService{
		baseService: newBaseService(requestId, db),
	}
}

func (userService UserService) Login(param *LoginParam) interface{} {
	// panic(errors.New(errors.ErrDatabase, ""))
	return nil
}
