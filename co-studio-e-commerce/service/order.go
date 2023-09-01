package service

import "co-studio-e-commerce/repo"

type Order struct {
	repo repo.IRepo
}

type IOrder interface {
}

func NewOrder(repo repo.IRepo) *Order {
	return &Order{repo: repo}
}
