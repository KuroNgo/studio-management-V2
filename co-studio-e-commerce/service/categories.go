package service

import (
	"co-studio-e-commerce/repo"
)

type CategoryService struct {
	repo repo.IRepo
}

type ICategoryService interface {
	// GetCategories() ([]*repo.Category, error)
}

func NewCategoryService(repo repo.IRepo) *CategoryService {
	return &CategoryService{repo: repo}
}
