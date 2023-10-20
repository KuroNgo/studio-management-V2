package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
)

type Category struct {
	repo     repo.IRepo
	category model.Categories
}

type ICategory interface {
	GetCategory(oneCategory model.Categories) (model.Categories, error)
	GetAllCategories() ([]model.Categories, error)
	CreateCategory(categoryRequest model.Categories) (model.Categories, error)
	DeleteCategoryFirst(oneCategory model.Categories) error
}

func NewCategory(repo repo.IRepo) *Category {
	return &Category{repo: repo}
}

// chhỉnh lại
func (c *Category) GetCategory(oneCategory model.Categories) (model.Categories, error) {
	err := c.repo.GetCategory(oneCategory)
	if err != nil {
		return model.Categories{}, err
	}
	return oneCategory, nil
}

// done
func (c *Category) GetAllCategories() ([]model.Categories, error) {
	// Gọi hàm GetAllCategory từ repo để lấy danh sách tất cả category
	categories, err := c.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) CreateCategory(categoryRequest model.Categories) (model.Categories, error) {
	category, err := c.repo.CreateCategory(categoryRequest)
	if err != nil {
		return model.Categories{}, err
	}
	return category, nil

}

func (c *Category) DeleteCategoryFirst(oneCategory model.Categories) error {
	category, err := c.repo.GetCategoryByID(oneCategory.ID)
	if err != nil {
		return err
	}

	err = c.repo.DisableCategory(&category)
	if err != nil {
		return err
	}
	return nil
}
