package service

import (
	"co-studio-e-commerce/common"
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
	"context"
)

type CategoryService struct {
	repo     repo.IRepo
	category model.Categories
}

type ICategoryService interface {
	GetCategory(oneCategory model.Categories) (model.Categories, error)
	GetAllCategories() ([]model.Categories, error)
	CreateCategory(oneCategory model.Categories) (categoriRes model.Categories, err error)
	UpdateCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
	DeleteCategory(ctx context.Context, oneCategory model.Categories) (categoriRes model.Categories, err error)
}

func NewCategoryService(repo repo.IRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

// chhỉnh lại
func (service *CategoryService) GetCategory(oneCategory model.Categories) (model.Categories, error) {
	err := service.repo.GetCategory(oneCategory)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}

// done
func (category *CategoryService) GetAllCategories() ([]model.Categories, error) {
	// Gọi hàm GetAllCategory từ repo để lấy danh sách tất cả category
	categories, err := category.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

//func (category *CategoryService) CreateCategory(categoryRequest model.Categories) (model.Categories, error) {
//	category, err := category.repo.CreateCategory(categoryRequest)
//
//}

func (service *CategoryService) DeleteCategory(oneCategory model.Categories) (categoriRes model.Categories, err error) {
	common.Sync(oneCategory, &categoriRes)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}
