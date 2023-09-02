package handler

import (
	"co-studio-e-commerce/service"
)

type Categorry struct {
	service service.ICategoryService
}

type ILogSession interface {
}

func NewLogSession(service service.ICategoryService) *Categorry {
	return &Categorry{service: service}
}

// hiển thị loại sản phẩm
// tất cả đều có theer thấy
// func (handler *Categorry) GetCategories(ctx *gin.Context, oneCategory model.Categories) (categoriRes model.Categories, err error) {
//
// }
