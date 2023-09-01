package service

import "co-studio-e-commerce/repo"

type Admin struct {
	repo repo.IRepo
}

type IAdmin interface {
}

func NewAdmin(repo repo.IRepo) *Admin {
	return &Admin{repo: repo}
}
