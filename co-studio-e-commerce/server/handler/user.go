package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
)

type User struct {
	service service.IUser
	user    model.User
}

type IUser interface {
}

func NewUser(service service.IUser) *User {
	return &User{service: service}
}
