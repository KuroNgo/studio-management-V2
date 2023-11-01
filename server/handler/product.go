package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
		MinPrice int32 `json:"minPrice"`
		MaxPrice int32 `json:"maxPrice"`
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
	var whoUpdate string
	if err := ctx.ShouldBindJSON(&whoUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	product, err := p.productService.GetProductByWhoUpdateForView(whoUpdate)
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

func (p *Product) GetProductByUpdateDateForView(ctx *gin.Context) {
	var updateDate time.Time
	if err := ctx.ShouldBindJSON(&updateDate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	product, err := p.productService.GetProductByUpdateDateForView(updateDate)
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

func (p *Product) CreateProduct(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser")
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := p.userService.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	product.WhoUpdate = userName.WhoUpdates
	data, err := p.productService.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
