package service

import (
	"github.com/chuxinplan/gin-mvc/app/controller"
)

type UserService struct {
	userId int64
}

func (userService UserService) Login(param *controller.LoginParam) {

}
