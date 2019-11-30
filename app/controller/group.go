package controller

import (
	"github.com/chuxinplan/gin-mvc/app/service"
	"github.com/gin-gonic/gin"
)

func HttpHandlerGetAllByUser(c *gin.Context) {
	groupService := service.NewGroupService(GetRequestId(c), GetDBSession(c), GetUserId(c))
	groups := groupService.GetAllGroupByUser()
	c.JSON(Success(groups))
}
