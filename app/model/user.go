package model

type User struct {
	Id       int64
	Email    string
	Username string
	Password string
}

func UserCreate(email string, username string, password string) (int64, error) {
	return 0, nil
}
