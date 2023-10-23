package service

import "co-studio-e-commerce/repo"

type Reviews struct {
	repo repo.IRepo
}

type IReviews interface {
}

func NewReviews(repo repo.IRepo) *Reviews {
	return &Reviews{repo: repo}
}
