package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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

// GetAllProductForView godoc
// @Summary get all products
// @Description get all products
// @Tags products
// @Accept json
// @Produce json
// @Router /client/product/get-all [get]
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

// GetProductByCategoryIDForView godoc
// @Summary get product by id
// @Description get product by id
// @Tags products
// @Accept json
// @Produce json
// @Router /client/product/get-id [get]
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

// GetProductByNameForView godoc
// @Summary get product by name
// @Description get product by name
// @Tags products
// @Accept json
// @Produce json
// @Router /client/product/get-name [get]
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

// GetProductByPriceForView godoc
// @Summary get product by price
// @Description get product by price
// @Tags products
// @Accept json
// @Produce json
// @Router /client/product/get-price [get]
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

// GetProductByWhoUpdateForView godoc
// @Summary get product by who updates
// @Description get product by who updates
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/get-who-update [get]
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

// GetProductByUpdateDateForView godoc
// @Summary get product by updates date
// @Description get product by updates date
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/get-update-date [get]
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

// CreateProduct godoc
// @Summary create product
// @Description create product
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/create [post]
func (p *Product) CreateProduct(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser")

	// Thực hiện lấy thông tin file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// Thực hiện kiểm tra có phải là file image
	if !util.IsImageFile(file.Filename) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This is not an image",
		})
		return
	}

	filePath := filepath.Join("uploads", file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

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

	product.AvatarURL = filePath
	product.WhoUpdates = userName.WhoUpdates
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

// UpdateProduct godoc
// @Summary update product
// @Description update product
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/update [post]
func (p *Product) UpdateProduct(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser")
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if !util.IsImageFile(file.Filename) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This is not an image",
		})
	}

	filePath := filepath.Join("uploads", file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
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

	data := model.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Description: product.Description,
		AvatarURL:   filePath,
		Enable:      product.Enable,
		IsUpdate:    1,
		WhoUpdates:  userName.WhoUpdates,
		UpdatedAt:   time.Now(),
		IsDelete:    1,
		CategoryID:  product.CategoryID,
	}

	result := p.productService.UpdateProduct(data)
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// UpdateEnable godoc
// @Summary update-enable product
// @Description update product
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/update-enable [patch]
func (p *Product) UpdateEnable(ctx *gin.Context) {
	var enable int
	if err := ctx.ShouldBindJSON(enable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	result := p.productService.UpdateEnable(enable)
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": result.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// Disable godoc
// @Summary update product
// @Description update product
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/remove-first [patch]
func (p *Product) Disable(ctx *gin.Context) {
	result := p.productService.Disable()
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": result.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// Enable godoc
// @Summary update enable product
// @Description update enable product
// @Tags products
// @Accept json
// @Produce json
// @Router /admin/product/enable [patch]
func (p *Product) Enable(ctx *gin.Context) {
	result := p.productService.Enable()
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": result.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
