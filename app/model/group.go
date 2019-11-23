package model

import "github.com/go-xorm/xorm"

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

func GroupCreate(db *xorm.Session) {

}

func GroupUpdate(db *xorm.Session) {

}

func GroupDelete(db *xorm.Session) {

}

func GroupGet(db *xorm.Session) {

}

func GroupGetBO(db *xorm.Session) {

}
