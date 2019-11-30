package model

import (
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
)

type Group struct {
	Id          int64
	Name        string
	Description string
	Users       string
}

type GroupBO struct {
	Id          int64
	Name        string
	Description string
	Users       []*User
}

func GroupCreate(db *xorm.Session, groupBO *GroupBO) (int64, error) {
	var userStr string
	for _, user := range groupBO.Users {
		userStr += strconv.FormatInt(user.Id, 10) + ","
	}
	userStr = strings.TrimRight(userStr, ",")
	group := &Group{
		Id:          groupBO.Id,
		Name:        groupBO.Name,
		Description: groupBO.Description,
		Users:       userStr,
	}
	_, err := db.Insert(group)
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func GroupUpdate(db *xorm.Session, groupBO *GroupBO) error {
	var userStr string
	for _, user := range groupBO.Users {
		userStr += strconv.FormatInt(user.Id, 10) + ","
	}
	userStr = strings.TrimRight(userStr, ",")
	group := &Group{
		Id:          groupBO.Id,
		Name:        groupBO.Name,
		Description: groupBO.Description,
		Users:       userStr,
	}
	_, err := db.AllCols().ID(groupBO.Id).Update(group)
	return err
}

func GroupDelete(db *xorm.Session, id int64) error {
	group := &Group{}
	_, err := db.ID(id).Delete(group)
	return err
}

func GroupGet(db *xorm.Session, id int64) (*Group, error) {
	group := &Group{}
	has, err := db.ID(id).Get(group)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return group, nil
}

func GetAllGroupByUser(db *xorm.Session, id int64) ([]*GroupBO, error) {
	var groupList []*Group
	idStr := strconv.FormatInt(id, 10)
	err := db.Where("users like ?", "%"+idStr+"%").Find(&groupList)
	if err != nil {
		return nil, err
	}
	if groupList == nil {
		return nil, nil
	}

	var groupBOList []*GroupBO
	for _, group := range groupList {
		userStrs := strings.Split(group.Users, ",")
		var userIds []int64
		for _, str := range userStrs {
			id, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			userIds = append(userIds, id)
		}
		var userList []*User
		if err := db.In("id", userIds).Find(&userList); err != nil {
			return nil, err
		}
		groupBO := &GroupBO{
			Id:          group.Id,
			Name:        group.Name,
			Description: group.Description,
			Users:       userList,
		}
		groupBOList = append(groupBOList, groupBO)
	}

	return groupBOList, nil
}

func GroupGetBO(db *xorm.Session, id int64) (*GroupBO, error) {
	group := &Group{}
	has, err := db.ID(id).Get(group)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	userStrs := strings.Split(group.Users, ",")
	var userIds []int64
	for _, str := range userStrs {
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, id)
	}
	var userList []*User
	if err := db.In("id", userIds).Find(&userList); err != nil {
		return nil, err
	}
	groupBO := &GroupBO{
		Id:          group.Id,
		Name:        group.Name,
		Description: group.Description,
		Users:       userList,
	}
	return groupBO, nil
}
