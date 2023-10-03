package repo

import (
	"errors"
	"github.com/google/uuid"
	"time"

	"co-studio-e-commerce/model"
)

// Hiển thị thông tin sản phẩm
func (r *Repo) GetProductsByCategory(categoryID uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("category_id = ?", categoryID.String()).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Hiển thị sản phẩm theo tên
func (r *Repo) GetProductByName(name string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("product_name = ?", name).Find(&products); err != nil {
		return nil, err.Error
	}
	return products, nil
}

// Hển t
func (r *Repo) GetProductByPrice(minPrice int32, maxPrice int32) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("price >= ? and price <= ?", minPrice, maxPrice).Find(&products); err != nil {
		return nil, errors.New("Find not Found")
	}
	return products, nil
}

func (r *Repo) GetProductByWho_Update(person string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("who_update = ?", person).Find(&products).Error; err != nil {
		return []model.Product{}, err
	}
	return products, nil

}

func (r *Repo) GetProductBy_UpdateDate(date time.Time) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where("update_date = ?", date).Find(&products).Error; err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (r *Repo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

// Tạo mới category
func (r *Repo) CreateProduct(product *model.Product) error {
	// CreateCategory là hàm tạo mới category
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateProduct(product *model.Product) error {
	if err := r.db.Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateEnable(enable int) error {
	var product *model.Product
	if err := r.db.Model(product).Update("enable = ?", enable).Error; err != nil {
		return err
	}
	if _enable := r.db.Model(product).Where("enable = 0").Update("is_delete = ?", 1).Error; _enable != nil {
		return _enable
	}
	return nil
}

func (r *Repo) Disable() error {
	var product *model.Product
	if err := r.db.Model(product).Update("enable = ? and is_delete = 0 ", 0).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) Enable() error {
	var product *model.Product
	if err := r.db.Model(product).Update("enable = ?", 1).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteProduct(product *model.Product) error {
	var count int64
	if err := r.db.Model(&model.Reviews{}).
		Where("product_id", product.ID).
		Count(&count).Error; err != nil {
		return err
	}

	if err := r.db.Model(&model.OrderDetail{}).
		Where("product_id", product.ID).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("Không thể xóa sản phẩm! Kiểm tra thông tin về reviews hoặc order_detail")
	}

	if err := r.db.Delete(product).Error; err != nil {
		return err
	}
	return nil
}
