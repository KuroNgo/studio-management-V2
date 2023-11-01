package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	productService service.IProduct
	userService    service.IUser
	productModel   model.Product
}

func NewProduct(product service.IProduct, user service.IUser) *Product {
	return &Product{
		productService: product,
		userService:    user,
	}
}

func (p *Product) GetAllProductForView(ctx *gin.Context) {
	product, err := p.productService.GetAllProductForView()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})

}

func (p *Product) GetProductByCategoryIDForView(ctx *gin.Context) {
	categoryID := ctx.Param("categoryid_id")
	product, err := p.productService.GetProductByCategoryForView(categoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})

}

func (p *Product) GetProductByNameForView(ctx *gin.Context) {
	nameProduct := ctx.Param("product_name")
	product, err := p.productService.GetProductByNameForView(nameProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (p *Product) GetProductByPriceForView(ctx *gin.Context) {
	var minPrice, maxPrice int32
	if err := ctx.ShouldBindJSON(&struct {
		minPrice int32 `json:"minPrice"`
		maxPrice int32 `json:"maxPrice"`
	}{minPrice, maxPrice}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	product, err := p.productService.GetProductByPriceForView(minPrice, maxPrice)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (p *Product) GetProductByWhoUpdateForView(ctx *gin.Context) {
	whoUpdate := ctx.Param("who_update")
	product, err := p.productService.GetProductByCategoryForView(whoUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}
