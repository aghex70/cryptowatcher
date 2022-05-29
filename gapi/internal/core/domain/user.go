package domain

type User struct {
	UserID     int64
	Source     string
	ExternalID int
}

func NewUser() User {
	return User{}
}
