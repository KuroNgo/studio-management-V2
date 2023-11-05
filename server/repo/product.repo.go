package repo

import (
	"errors"
	"gorm.io/gorm"
	"time"

	"co-studio-e-commerce/model"
)

// GetProductsByCategory Hiển thị thông tin sản phẩm
func (r *Repo) GetProductsByCategory(categoryID string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Where("category_id = ?", categoryID).
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("product not found")
		}
	}
	return products, nil
}

// GetProductByName Hiển thị sản phẩm theo tên
func (r *Repo) GetProductByName(name string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Where("product_name = ?", name).
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("product not found")
		}
	}
	return products, nil
}

func (r *Repo) GetProductByPrice(minPrice int32, maxPrice int32) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Where("min_price >= ? and max_price <= ?", minPrice, maxPrice).
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("product not found")
		}
	}
	return products, nil
}

func (r *Repo) GetProductByWhoUpdate(person string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Where("who_update = ?", person).
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("product not found")
		}
	}
	return products, nil
}

func (r *Repo) GetProductByUpdateDate(date time.Time) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Where("update_date = ?", date).
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("product not found")
		}
	}
	return products, nil
}

func (r *Repo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		Find(&products).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Product{}, errors.New("no product created")
		}
	}
	return products, nil
}

// CreateProduct Tạo mới category
func (r *Repo) CreateProduct(product model.Product) (model.Product, error) {
	// CreateCategory là hàm tạo mới category
	if err := r.db.
		Create(product).
		Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *Repo) UpdateProduct(product *model.Product) error {
	if err := r.db.
		Updates(product).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateEnable(enable int) error {
	var product *model.Product
	if err := r.db.
		Where("product_id = ?", product.ID).
		Update("enable = ?", enable).
		Error; err != nil {
		return err
	}
	if _enable := r.db.
		Where("enable = 0").
		Update("is_delete = ?", 1).
		Error; _enable != nil {
		return _enable
	}
	return nil
}

func (r *Repo) Disable() error {
	var product *model.Product
	if err := r.db.
		Where("product_id = ?", product.ID).
		Update("enable = ? and is_delete = 1 ", 0).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) Enable() error {
	var product *model.Product
	if err := r.db.
		Where("product_id = ?", product.ID).
		Update("enable = ? and is_delete = 0", 1).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteProduct(product *model.Product) error {
	var count int64
	if err := r.db.Model(&model.Reviews{}).
		Where("product_id", product.ID).
		Count(&count).
		Error; err != nil {
		return err
	}

	if err := r.db.Model(&model.OrderDetail{}).
		Where("product_id", product.ID).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("không thể xóa sản phẩm! Kiểm tra thông tin về reviews hoặc order_detail")
	}

	if err := r.db.Delete(product).Error; err != nil {
		return err
	}
	return nil
}
