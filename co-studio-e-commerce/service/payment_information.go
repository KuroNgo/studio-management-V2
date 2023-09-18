package service

import "co-studio-e-commerce/repo"

type PaymentInformation struct {
	repo repo.IRepo
}

type IPaymentInformation interface {
}

func NewPaymentInformation(repo repo.IRepo) *PaymentInformation {
	return &PaymentInformation{repo: repo}
}
