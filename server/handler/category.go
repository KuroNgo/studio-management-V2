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

// GetAllCategoriesForView godoc
// @Summary get all categories
// @Description get all categories
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/get-all [get]
func (c *Category) GetAllCategoriesForView(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"category": categories,
		},
	})
}

// GetCategoryByIDForView godoc
// @Summary get category by id
// @Description get category by id
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/get/:category_id [get]
func (c *Category) GetCategoryByIDForView(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")
	category, err := c.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

// GetCategoryByIDForEdit godoc
// @Summary get category by id
// @Description get category by id
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/get2/:category_id [get]
func (c *Category) GetCategoryByIDForEdit(ctx *gin.Context) {
	categoryID := ctx.Param("category_id")
	category, err := c.categoryService.GetCategoryByID2(categoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

// GetCategoryByUpdateDateForEdit godoc
// @Summary get category by update date
// @Description get category by update date
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/get-update-date/:update_date [get]
func (c *Category) GetCategoryByUpdateDateForEdit(ctx *gin.Context) {
	updateDateParam := ctx.Param("update_date")
	updateDateParam2, err := time.Parse("02-01-2006", updateDateParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'date' parameter",
		})
		return
	}
	categories, err := c.categoryService.GetCategoryCreateByUpdateDate(updateDateParam2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Trả về mảng dữ liệu dưới dạng JSON
	ctx.JSON(http.StatusOK, categories)
}

// GetCategoryByEnableForEdit godoc
// @Summary get category by update date
// @Description get category by update date
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/get-enable/:enable [get]
func (c *Category) GetCategoryByEnableForEdit(ctx *gin.Context) {
	enableParam := ctx.Param("enable")
	enable, err := strconv.Atoi(enableParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'enable' parameter",
		})
		return
	}
	categories, err := c.categoryService.GetCategoryByEnable(enable)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary create category
// @Description create category
// @Tags category
// @Accept json
// @Produce json
// @Router /admin/category/create [post]
func (c *Category) CreateCategory(ctx *gin.Context) {
	// Lấy user hiện đang đăng nhập
	currentUser := ctx.MustGet("currentUser")

	// Thực hiện tạo input cho người dùng
	var category model.CategoryForCreate
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

	// Thực hiện thêm category
	newCategory := model.Category{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       1,
		IsUpdate:     0,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
		WhoUpdate:    userName.FullName,
		IsDelete:     0,
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

// UpdateCategory godoc
// @Summary update category
// @Description create category
// @Tags category
// @Accept json
// @Produce json
// @Router /admin/category/update [put]
func (c *Category) UpdateCategory(ctx *gin.Context) {
	// Lấy user hiện tại đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.CategoryForUpdate
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

	updateCategory := model.Category{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       category.Enable,
		IsUpdate:     1,
		WhoUpdate:    userName.FullName,
		UpdatedAt:    time.Now(),
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

// DeleteCategoryFirst godoc
// @Summary delete category
// @Description delete category
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/delete [patch]
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

// DeleteCategorySecond godoc
// @Summary delete category with admin
// @Description delete category with admin
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/delete-second [delete]
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

// ResolveCategory godoc
// @Summary resolve category with admin
// @Description resolve category with admin
// @Tags category
// @Accept json
// @Produce json
// @Router /client/category/resolve [patch]
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
