package service

import "github.com/go-xorm/xorm"

type GroupService struct {
	baseService
	Username string
}

func NewGroupService(requestId string, db *xorm.Session, username string) GroupService {
	return GroupService{
		baseService: newBaseService(requestId, db),
		Username:    username,
	}
}

func (groupService GroupService) GetAllGroupByUser() {

}
