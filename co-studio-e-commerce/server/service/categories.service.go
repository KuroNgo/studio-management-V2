package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
)

type Category struct {
	repo     repo.IRepo
	category model.Categories
	user     model.User
}

type ICategories interface {
	GetCategory(uuid string) (model.Categories, error)
	GetAllCategories() ([]model.Categories, error)
	CreateCategory(categoryRequest model.Categories) (model.Categories, error)
	UpdateCategory(categoryRequest model.Categories) (model.Categories, error)
	DeleteCategoryFirst(category model.Categories) error
	DeleteCategorySecond(category model.Categories) error
}

func NewCategory(repo repo.IRepo) *Category {
	return &Category{repo: repo}
}

func (c *Category) GetCategory(uuid string) (model.Categories, error) {
	category, err := c.repo.GetCategoryByID(uuid)
	if err != nil {
		return model.Categories{}, err
	}
	return category, nil
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

func (c *Category) UpdateCategory(categoryRequest model.Categories) (model.Categories, error) {
	category, err := c.repo.UpdateCategory(&categoryRequest)
	if err != nil {
		return model.Categories{}, err
	}
	return category, nil
}

func (c *Category) DeleteCategoryFirst(category model.Categories) error {
	err := c.repo.DisableCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) DeleteCategorySecond(category model.Categories) error {
	err := c.repo.DeleteCategory(&category)
	if err != nil {
		return err
	}
	return nil
}
