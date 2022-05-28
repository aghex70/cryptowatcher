package domain

type User struct {
	UserID int64
}

func NewUser() User {
	return User{}
}
