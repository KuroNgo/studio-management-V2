package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
	"time"
)

type Category struct {
	repo     repo.IRepo
	category model.Category
}

type ICategory interface {
	GetCategoryByID(uuid string) (model.Category, error)
	GetCategoryByID2(uuid string) (*model.Category, error)
	GetCategoryByEnable(enable int) (*[]model.Category, error)
	GetCategoryByName(name string) (*[]model.Category, error)
	GetAllCategories() ([]model.Category, error)
	GetCategoryCreateByUpdateDate(date time.Time) (*[]model.Category, error)
	CreateCategory(categoryRequest model.Category) (model.Category, error)
	UpdateCategory(categoryRequest model.Category) (model.Category, error)
	EnableCategory(category model.Category) error
	DeleteCategoryFirst(category model.Category) error
	DeleteCategorySecond(category model.Category) error
}

func NewCategory(repo repo.IRepo) *Category {
	return &Category{repo: repo}
}

// GetCategoryByID Nếu chỉ get ra để xem
func (c *Category) GetCategoryByID(uuid string) (model.Category, error) {
	category, err := c.repo.GetCategoryByID(uuid)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// GetCategoryByID2 Nếu get ra để edit
func (c *Category) GetCategoryByID2(uuid string) (*model.Category, error) {
	category, err := c.repo.GetCategoryByID2(uuid)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) GetCategoryByEnable(enable int) (*[]model.Category, error) {
	category, err := c.repo.GetCategoryByEnable(enable)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetCategoryByName Nếu get ra để xử lý
func (c *Category) GetCategoryByName(name string) (*[]model.Category, error) {
	category, err := c.repo.GetCategoryByName(name)
	if err != nil {
		return &[]model.Category{}, err
	}
	return category, nil
}

// GetAllCategories done
func (c *Category) GetAllCategories() ([]model.Category, error) {
	// Gọi hàm GetAllCategory từ repo để lấy danh sách tất cả category
	categories, err := c.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetCategoryCreateByUpdateDate(date time.Time) (*[]model.Category, error) {
	category, err := c.repo.GetCategoryCreatedByUpdateDate(date)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) CreateCategory(categoryRequest model.Category) (model.Category, error) {
	category, err := c.repo.CreateCategory(categoryRequest)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil

}

func (c *Category) UpdateCategory(categoryRequest model.Category) (model.Category, error) {
	category, err := c.repo.UpdateCategory(&categoryRequest)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (c *Category) EnableCategory(category model.Category) error {
	err := c.repo.EnableCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) DeleteCategoryFirst(category model.Category) error {
	err := c.repo.DisableCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) DeleteCategorySecond(category model.Category) error {
	err := c.repo.DeleteCategory(&category)
	if err != nil {
		return err
	}
	return nil
}
