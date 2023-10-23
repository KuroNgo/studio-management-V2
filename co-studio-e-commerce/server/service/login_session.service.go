package service

import "co-studio-e-commerce/repo"

type LoginSession struct {
	repo repo.IRepo
}

type ILoginSession interface {
}

func NewLoginSession(repo repo.IRepo) *LoginSession {
	return &LoginSession{repo: repo}
}
