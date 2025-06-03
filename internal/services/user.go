package services

import "api/internal/repositories"

type User struct {
	repoisitory repositories.User
}

func NewUser() *User {
	return &User{
		repoisitory: *repositories.NewUser(),
	}
}
