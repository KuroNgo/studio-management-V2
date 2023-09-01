package service

import "co-studio-e-commerce/repo"

type ShoppingCart struct {
	repo repo.IRepo
}

type IShoppingCart interface {
}

func NewShoppingCart(repo repo.IRepo) *ShoppingCart {
	return &ShoppingCart{repo: repo}
}
