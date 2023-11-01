package repo

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"

	"co-studio-e-commerce/model"
)

func (r *Repo) GetOrderByID(id int) (model.Order, error) {
	// GetOrder là hàm lấy thông tin order
	var order model.Order
	if err := r.db.Where("id = ?", id).First(order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByOrderDate(date time.Time) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("order_date = ?", date).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByOrderDateAndTotalAmount(orderDate time.Time, totalAmount int32) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("order_date = ? and total_amount = ?", orderDate, totalAmount).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByTotalAmount(totalAmount int32) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("total_amount = ?", totalAmount).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetAllOrder() ([]model.Order, error) {
	var order []model.Order
	// GetAllOrder là hàm lấy thông tin tất cả order
	if err := r.db.Find(order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByUserID(uuid uuid.UUID) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("user_id = ?", uuid.String()).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByEnable(enable int) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("enable = ?", enable).Find(&order).Limit(5).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) GetOrderByIsUpdate(isUpdate int) ([]model.Order, error) {
	var order []model.Order
	if err := r.db.Where("is_update = ?", isUpdate).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Order{}, nil
		}
	}
	return order, nil
}

func (r *Repo) CreateOrder(order *model.Order) error {
	// CreateOrder là hàm tạo mới order
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateOrder(order *model.Order) error {
	// UpdateOrder là hàm cập nhật thông tin order
	if err := r.db.Save(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteOrder(order *model.Order) error {
	// DeleteOrder là hàm xóa thông tin order
	if err := r.db.Delete(order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) RemoveOrder(order *model.Order) error {
	// RemoveOrder là hàm xóa thông tin order
	// Unscoped() để xóa cả dữ liệu đã bị soft delete
	// soft delete là xóa dữ liệu nhưng không xóa vĩnh viễn, chỉ đánh dấu là đã xóa
	if err := r.db.Unscoped().Delete(order).Error; err != nil {
		return err
	}
	return nil
}
