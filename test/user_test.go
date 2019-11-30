package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/chuxinplan/gin-mvc/common/db"
	"github.com/chuxinplan/gin-mvc/common/errors"

	"github.com/chuxinplan/gin-mvc/app/model"

	"github.com/chuxinplan/gin-mvc/router"
	"gopkg.in/go-playground/assert.v1"
)

var (
	user = &model.User{
		Id:       0,
		Email:    "test@gmail.com",
		Username: "test123",
		Password: "Test123!",
	}
)

func TestUserRegisterSucc(t *testing.T) {
	session := db.DB.NewSession()
	userInfo, err := model.UserGetByUsername(session, user.Username)
	if err != nil {
		panic(errors.Warp(errors.ErrDatabase, err.Error()))
	}

	if userInfo != nil {
		err = model.UserDelete(session, userInfo.Id)
		if err != nil {
			panic(errors.Warp(errors.ErrDatabase, err.Error()))
		}
	}

	data := url.Values{"email": {user.Email}, "username": {user.Username}, "password": {user.Password}}
	body := strings.NewReader(data.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/register", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"code\":0,\"message\":\"success\",\"data\":null}", w.Body.String())
}

func TestUserLoginSucc(t *testing.T) {
	data := url.Values{"type": {"username"}, "account": {user.Username}, "password": {user.Password}}
	body := strings.NewReader(data.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/login", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
