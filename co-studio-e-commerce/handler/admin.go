package handler

import "co-studio-e-commerce/service"

type Admin struct {
	service service.IAdmin
}

type IAdmin interface {
}

func NewAdmin(service service.IAdmin) *Admin {
	return &Admin{service: service}
}

func (handler *Admin) GetAdmins() {

}
