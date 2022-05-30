package domain

type User struct {
	ID         uint
	ExternalID int
	Source     string
}

func NewUser() User {
	return User{}
}
