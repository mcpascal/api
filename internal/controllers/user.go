package controllers

import "api/internal/services"

type User struct {
	service services.User
}

func NewUser(service services.User) *User {
	return &User{
		service: service,
	}
}
