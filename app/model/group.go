package model

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
