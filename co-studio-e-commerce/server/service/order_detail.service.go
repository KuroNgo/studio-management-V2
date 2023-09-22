package service

import "co-studio-e-commerce/repo"

type OrderDetail struct {
	repo repo.IRepo
}

type IOrderDetail interface {
}

func NewOrderDetail(repo repo.IRepo) *OrderDetail {
	return &OrderDetail{repo: repo}
}
