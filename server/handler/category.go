package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Category struct {
	categoryService service.ICategory
	userService     service.IUser
	categoryModel   model.Category
}

func NewCategory(service service.ICategory, user service.IUser) *Category {
	return &Category{
		categoryService: service,
		userService:     user,
	}
}

// GetAllCategoriesForView GetAllCategories Lấy tất cả loại
func (c *Category) GetAllCategoriesForView(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"category": categories}})
}

// GetCategoryByIDForView GetCategoryByID tìm category theo id nếu để xem
func (c *Category) GetCategoryByIDForView(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")
	category, err := c.categoryService.GetCategoryByID(categoryID)
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

// GetCategoryByIDForEdit GetCategoryByID tìm category theo id nếu để edit
func (c *Category) GetCategoryByIDForEdit(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")
	category, err := c.categoryService.GetCategoryByID2(categoryID)
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

func (c *Category) GetCategoryByUpdateDateForEdit(ctx *gin.Context) {
	updateDateParam := ctx.Param("update_date")
	updateDateParam2, err := time.Parse("02-01-2006", updateDateParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'date' parameter"})
		return
	}
	categories, err := c.categoryService.GetCategoryCreateByUpdateDate(updateDateParam2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về mảng dữ liệu dưới dạng JSON
	ctx.JSON(http.StatusOK, categories)
}

// GetCategoryByEnableForEdit lấy category theo enable
func (c *Category) GetCategoryByEnableForEdit(ctx *gin.Context) {
	enableParam := ctx.Param("enable")
	enable, err := strconv.Atoi(enableParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'enable' parameter"})
		return
	}
	categories, err := c.categoryService.GetCategoryByEnable(enable)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// CreateCategory Tạo mới category
func (c *Category) CreateCategory(ctx *gin.Context) {
	// Lấy user hiện đang đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := c.userService.FindUserByID(fmt.Sprint(currentUser))
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
	newCategory := model.Category{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       1,
		IsUpdate:     category.IsUpdate,
		UpdatedAt:    category.UpdatedAt,
		//WhoUpdate:    fmt.Sprint(currentUser),
		WhoUpdate: userName.FullName,
		IsDelete:  0,
	}

	categoryResponse, err := c.categoryService.CreateCategory(newCategory)
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

// UpdateCategory Thực hiện cập nhật category
func (c *Category) UpdateCategory(ctx *gin.Context) {
	// Lấy user hiện tại đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// Thực hiện lấy thông tin người dùng
	userName, err := c.userService.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.UpdatedAt = time.Now()
	updateCategory := model.Category{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       category.Enable,
		IsUpdate:     1,
		WhoUpdate:    userName.FullName,
		UpdatedAt:    category.UpdatedAt,
		IsDelete:     0,
	}

	data, err := c.categoryService.UpdateCategory(updateCategory)
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

// DeleteCategoryFirst Thực hiện chuyển enable về disable thông qua câu lệnh update
func (c *Category) DeleteCategoryFirst(ctx *gin.Context) {
	// Lấy user hiện tại đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := c.userService.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.UpdatedAt = time.Now()
	deleteCategory := model.Category{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       0,
		IsUpdate:     1,
		UpdatedAt:    category.UpdatedAt,
		WhoUpdate:    userName.FullName,
		IsDelete:     1,
	}

	err = c.categoryService.DeleteCategoryFirst(deleteCategory)
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

// DeleteCategorySecond Chỉ admin mới xóa được lần 2
func (c *Category) DeleteCategorySecond(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser")
	var category model.Category

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

	err := c.categoryService.DeleteCategorySecond(category)
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

// ResolveCategory Khôi phục category đã xóa ( disable to enable)
func (c *Category) ResolveCategory(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser")
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	//Thực hiện lấy thông tin người dùng
	userName, err := c.userService.FindUserByID(fmt.Sprint(currentUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	category.WhoUpdate = userName.FullName
	err = c.categoryService.EnableCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
