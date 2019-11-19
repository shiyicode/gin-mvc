package service

type GroupService struct {
	baseService
}

func NewGroupService(username string, requestId string) GroupService {
	return GroupService{
		baseService: newBaseService(username, requestId),
	}
}

func (groupService GroupService) GetAllGroup(param *LoginParam) {

}
