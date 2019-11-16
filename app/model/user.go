package model

type User struct {
	Id       int64
	Email    string
	Password string
	Username string
	Nickname string
}

type Book struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
}
