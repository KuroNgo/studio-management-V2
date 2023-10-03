package service

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"co-studio-e-commerce/common"
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
)

type CategoryService struct {
	repo repo.IRepo
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

func (service *CategoryService) GetCategory(oneCategory model.Categories) (model.Categories, error) {
	err := service.repo.GetCategory(&oneCategory)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}

func (service *CategoryService) GetAllCategories() ([]model.Categories, error) {
	// Gọi hàm GetAllCategory từ repo để lấy danh sách tất cả category
	categories, err := service.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (service *CategoryService) CreateCategory(category model.Categories) (model.Categories, error) {
	// Tạo một UUID mới cho CategoryID
	category.CategoryID = uuid.New()

	// Thử tạo mới danh mục
	createdCategory, err := service.repo.CreateCategory(category)
	if err != nil {
		return model.Categories{}, err
	}

	// Kiểm tra xem createdCategory.CategoryID có được thiết lập sau khi tạo không
	if createdCategory.CategoryID == uuid.Nil {
		return model.Categories{}, errors.New("Lỗi: category.CategoryID không được thiết lập sau khi tạo mới")
	}

	// Trả về danh mục đã được tạo mới
	return createdCategory, nil
}

func (service *CategoryService) DeleteCategory(oneCategory model.Categories) (categoriRes model.Categories, err error) {
	common.Sync(oneCategory, &categoriRes)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}
