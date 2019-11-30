package service

import (
	"github.com/chuxinplan/gin-mvc/app/model"
	"github.com/chuxinplan/gin-mvc/common/errors"
	"github.com/go-xorm/xorm"
)

type GroupService struct {
	baseService
	UserId int64
}

func NewGroupService(requestId string, db *xorm.Session, userId int64) GroupService {
	return GroupService{
		baseService: newBaseService(requestId, db),
		UserId:      userId,
	}
}

func (groupService GroupService) GetAllGroupByUser() []*model.GroupBO {
	groups, err := model.GetAllGroupByUser(groupService.DB, groupService.UserId)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	return groups
}
