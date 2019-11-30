package model

import (
	"github.com/go-xorm/xorm"
)

type User struct {
	Id       int64
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func UserCreate(db *xorm.Session, userInfo *User) (int64, error) {
	_, err := db.Insert(userInfo)
	if err != nil {
		return 0, err
	}
	return userInfo.Id, nil
}

func UserUpdate(db *xorm.Session, userInfo *User) error {
	_, err := db.AllCols().ID(userInfo.Id).Update(userInfo)
	return err
}

func UserDelete(db *xorm.Session, userId int64) error {
	user := &User{}
	_, err := db.ID(userId).Delete(user)
	return err
}

func UserGetById(db *xorm.Session, userId int64) (*User, error) {
	user := &User{}
	has, err := db.Where("id=?", userId).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func UserGetByUsername(db *xorm.Session, username string) (*User, error) {
	user := &User{}
	has, err := db.Where("username=?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func UserGetByEmail(db *xorm.Session, email string) (*User, error) {
	user := &User{}
	has, err := db.Where("email=?", email).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}
