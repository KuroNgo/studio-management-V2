package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Category struct {
	user     service.IUser
	service  service.ICategories
	category model.Categories
}

func NewCategory(service service.ICategories, user service.IUser) *Category {
	return &Category{
		service: service,
		user:    user,
	}
}

// Lấy tất cả loại
func (c *Category) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"category": categories}})
}

// tìm category theo id
func (c *Category) GetCategoryByID(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")
	category, err := c.service.GetCategory(categoryID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   category,
	})
}

// Tạo mới category
func (c *Category) CreateCategory(ctx *gin.Context) {
	// Lấy user hiện đang đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Categories
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := c.user.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.UpdatedAt = time.Now()
	category.CreatedAt = time.Now()

	// Thực hiện thêm category
	newCategory := model.Categories{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       1,
		IsUpdate:     category.IsUpdate,
		UpdatedAt:    category.UpdatedAt,
		//WhoUpdate:    fmt.Sprint(currentUser),
		WhoUpdate: userName.FullName,
		IsDelete:  0,
	}

	categoryResponse, err := c.service.CreateCategory(newCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   categoryResponse,
	})

}

// Thực hiện cập nhật category
func (c *Category) UpdateCategory(ctx *gin.Context) {
	// Lấy user hiện tại đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Categories
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// Thực hiện lấy thông tin người dùng
	userName, err := c.user.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.UpdatedAt = time.Now()
	updateCategory := model.Categories{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       category.Enable,
		IsUpdate:     1,
		WhoUpdate:    userName.FullName,
		UpdatedAt:    category.UpdatedAt,
		IsDelete:     0,
	}

	data, err := c.service.UpdateCategory(updateCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Thực hiện chuyển enable về disable thông qua câu lệnh update
func (c *Category) DeleteCategoryFirst(ctx *gin.Context) {
	// Lấy user hiện tại đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Categories
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := c.user.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.UpdatedAt = time.Now()
	deleteCategory := model.Categories{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       0,
		IsUpdate:     1,
		UpdatedAt:    category.UpdatedAt,
		WhoUpdate:    userName.FullName,
		IsDelete:     1,
	}

	err = c.service.DeleteCategoryFirst(deleteCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// Chỉ admin mới xóa được lần 2
func (c *Category) DeleteCategorySecond(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser")
	var category model.Categories

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// Thực hiện lấy thông tin người dùng
	//username, err := c.user.FindUserByID(fmt.Sprint(&currentUser))
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, gin.H{
	//		"status":  "fail",
	//		"message": err.Error(),
	//	})
	//	return
	//}

	err := c.service.DeleteCategorySecond(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

// Khôi phục category đã xóa ( disable to enable)
//func (c *Category) ResolveCategory(ctx *gin.Context) {
//	currentUser := ctx.MustGet("currentUser")
//	var category model.Categories
//	if err := ctx.ShouldBindJSON(&category); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "fail",
//			"message": err.Error(),
//		})
//		return
//	}
//
//	data, err := c.service.
//}
