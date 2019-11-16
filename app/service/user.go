package service

import (
	"github.com/chuxinplan/gin-mvc/app/controller"
	"github.com/chuxinplan/gin-mvc/common/errors"
)

type UserService struct {
	userId int64
}

func (userService UserService) Login(param *controller.LoginParam) (interface{}, *errors.Err) {
	return nil, errors.New(errors.ErrDatabase, "")
}
