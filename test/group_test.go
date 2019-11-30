package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chuxinplan/gin-mvc/common/auth"

	"github.com/chuxinplan/gin-mvc/app/model"
	"github.com/chuxinplan/gin-mvc/common/errors"

	"github.com/chuxinplan/gin-mvc/common/db"

	"github.com/chuxinplan/gin-mvc/router"
	assert "gopkg.in/go-playground/assert.v1"
)

// TODO group test
// TODO gourp 鉴权

var (
	userId    = int64(10001)
	gUserInfo = &model.User{
		Id:       10001,
		Email:    "test1@gmail.com",
		Username: "test1231",
		Password: "Test1231!",
	}
	group = &model.GroupBO{
		Id:          10010,
		Name:        "testGourp",
		Description: "testGroupDesc",
		Users:       nil,
	}
)

func TestGroupGetAllGroupByUserSucc(t *testing.T) {
	session := db.DB.NewSession()
	groupList, err := model.GetAllGroupByUser(session, userId)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	for i := 0; i < len(groupList); i++ {
		err = model.GroupDelete(session, groupList[i].Id)
		if err != nil {
			panic(errors.Warp(errors.ErrDatabase, err.Error()))
		}
	}

	var userList []*model.User
	userList = append(userList, gUserInfo)
	group.Users = userList
	err = model.UserDelete(session, gUserInfo.Id)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	_, err = model.UserCreate(session, gUserInfo)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}

	err = model.GroupDelete(session, group.Id)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	_, err = model.GroupCreate(session, group)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/auth/groups", strings.NewReader(""))

	token, err := auth.EncodeToken(gUserInfo.Username, gUserInfo.Id)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}
	cookie1 := &http.Cookie{Name: "token", Value: token, HttpOnly: true}
	req.AddCookie(cookie1)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"code\":0,\"message\":\"success\",\"data\":[{\"Id\":10010,\"Name\":\"testGourp\",\"Description\":\"testGroupDesc\",\"Users\":[{\"Id\":10001,\"email\":\"test1@gmail.com\",\"username\":\"test1231\",\"password\":\"Test1231!\"}]}]}", w.Body.String())
}
