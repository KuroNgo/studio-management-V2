package service

import "co-studio-e-commerce/repo"

type Product struct {
	repo repo.IRepo
}

type IProduct interface {
}

func NewProduct(repo repo.IRepo) *Product {
	return &Product{repo: repo}
}
