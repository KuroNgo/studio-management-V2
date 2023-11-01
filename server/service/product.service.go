package service

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/repo"
	"time"
)

type Product struct {
	repo    repo.IRepo
	product model.Product
}

type IProduct interface {
	GetProductByCategoryForView(categoryID string) ([]model.Product, error)
	GetProductByNameForView(name string) ([]model.Product, error)
	GetProductByPriceForView(minPrice int32, maxPrice int32) ([]model.Product, error)
	GetProductByWhoUpdateForView(person string) ([]model.Product, error)
	GetProductByUpdateDateForView(date time.Time) ([]model.Product, error)
	GetAllProductForView() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	UpdateProduct(product model.Product) error
	UpdateEnable(enable int) error
	Disable() error
	Enable() error
	DeleteProduct(product *model.Product) error
}

func NewProduct(repo repo.IRepo) *Product {
	return &Product{repo: repo}
}

func (p *Product) GetProductByCategoryForView(categoryID string) ([]model.Product, error) {
	product, err := p.repo.GetProductsByCategory(categoryID)
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (p *Product) GetProductByNameForView(name string) ([]model.Product, error) {
	product, err := p.repo.GetProductByName(name)
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (p *Product) GetProductByPriceForView(minPrice int32, maxPrice int32) ([]model.Product, error) {
	product, err := p.repo.GetProductByPrice(minPrice, maxPrice)
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (p *Product) GetProductByWhoUpdateForView(person string) ([]model.Product, error) {
	product, err := p.repo.GetProductByWhoUpdate(person)
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (p *Product) GetProductByUpdateDateForView(date time.Time) ([]model.Product, error) {
	product, err := p.repo.GetProductByUpdateDate(date)
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (p *Product) GetAllProductForView() ([]model.Product, error) {
	product, err := p.repo.GetAllProduct()
	if err != nil {
		return []model.Product{}, err
	}
	return product, err
}

func (p *Product) CreateProduct(product model.Product) (model.Product, error) {
	data, err := p.repo.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	return data, nil
}

func (p *Product) UpdateProduct(product model.Product) error {
	err := p.repo.UpdateProduct(&product)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateEnable(enable int) error {
	err := p.repo.UpdateEnable(enable)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Disable() error {
	err := p.repo.Disable()
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Enable() error {
	err := p.repo.Enable()
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) DeleteProduct(product *model.Product) error {
	err := p.repo.DeleteProduct(product)
	if err != nil {
		return err
	}
	return nil
}
