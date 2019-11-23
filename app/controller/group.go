package controller

import (
	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/gin-gonic/gin"
)

func HttpHandlerGetAllByUser(c *gin.Context) {
	//param := new(service.LoginParam)
	//if err := c.ShouldBind(param); err != nil {
	//	panic(errors.Warp(errors.ErrValidation, err.Error()))
	//}

	groupService := service.NewGroupService(getRequestId(c), getDBSession(c), getUsername(c))

	groupService.GetAllGroupByUser()

	c.JSON(Success())
}
