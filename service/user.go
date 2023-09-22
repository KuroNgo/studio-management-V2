package service

import "co-studio-e-commerce/repo"

type User struct {
	repo repo.IRepo
}

type IUser interface {
}

func NewUser(repo repo.IRepo) *User {
	return &User{repo: repo}
}
