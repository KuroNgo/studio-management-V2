package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Categorry struct {
	service  service.ICategoryService
	category model.Categories
}

func NewCategory(service service.ICategoryService) *Categorry {
	return &Categorry{service: service}
}

func (c *Categorry) GetAllCategories(ctx *gin.Context) {
	categorys, err := c.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"category": categorys}})
}

func (c *Categorry) GetCategory(category model.Categories) {

}
