package service

import (
	"context"

	"co-studio-e-commerce/common"
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
)

type CategoryService struct {
	repo repo.IRepo
}

type ICategoryService interface {
	GetCategories(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
	GetAllCategories(ctx context.Context, allCategory []model.Categories) (categoriRes []model.Categories, err error)
	CreateCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
	UpdateCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
	DeleteCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
}

func NewCategoryService(repo repo.IRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (service *CategoryService) GetCategories(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error) {
	err = service.repo.GetCategory(ctx, &oneCategory)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}

func (service *CategoryService) GetAllCategories(ctx context.Context, allCategory []model.Categories) (categoriRes []model.Categories, err error) {
	err = service.repo.GetAllCategory(ctx, &allCategory)
	if err != nil {
		return []model.Categories{}, err
	}
	return allCategory, nil
}

//
// func (service *CategoryService) CreateCategory(ctx context.Context, oneCategoryReq model.Categories) (categoriRes model.Categories, err error) {
// 	// Sử dụng common.Sync để sao chép dữ liệu từ oneCategoryReq sang categoriRes
// 	common.Sync(oneCategoryReq, &categoriRes)
//
// 	// Kiểm tra xem oneCategoryReq.CategoryName đã tồn tại trong danh sách
// 	if categoryExists(oneCategoryReq.CategoryName) {
// 		return model.Categories{}, errors.New("Category with the same name already exists")
// 	}
//
// 	// Gọi service.repo.CreateCategory để tạo danh mục mới
// 	if err := service.repo.CreateCategory(ctx, &categoriRes); err != nil {
// 		return model.Categories{}, err
// 	}
//
// 	return categoriRes, nil
// }

// func categoryExists(categoryName string) bool {
// 	// Thực hiện truy vấn SQL hoặc kiểm tra trong cơ sở dữ liệu
// 	// Nếu tìm thấy, trả về true, ngược lại trả về false
// 	// Đây chỉ là một ví dụ giả định
// 	// Ví dụ: Kiểm tra tên danh mục trong bảng 'categories'
// 	var count int
// 	db.Model(&model.Categories{}).Where("category_name = ?", categoryName).Count(&count)
// 	return count > 0
// }

// func (service *CategoryService) UpdateCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error) {
// 	common.Sync(oneCategory, &categoriRes)
// 	if err != nil {
// 		return model.Categories{}, err
// 	}
// 	// Kiểm tra xem oneCategory.CategoryName đã tồn tại trong danh sách chưa
// 	// Đây là một ví dụ giả định, bạn cần thay thế bằng kiểm tra thực tế trong danh sách
// 	if categoryExists(oneCategory.CategoryName) {
// 		return model.Categories{}, err
// 	}
// 	return oneCategory, nil
// }

func (service *CategoryService) DeleteCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error) {
	common.Sync(oneCategory, &categoriRes)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}
