package repo

import (
	"errors"
	"time"

	"co-studio-e-commerce/model"
)

// Nếu get để xem
func (r *Repo) GetCategoryByID(uuid string) (model.Categories, error) {
	var category model.Categories
	if err := r.db.
		Model(&category).
		Where("category_id = ? and enable = 1", uuid).
		First(&category).
		Error; err != nil {
		return model.Categories{}, err
	}
	return category, nil
}

// Nếu get để làm việc
func (r *Repo) GetCategoryByID2(uuid string) (*model.Categories, error) {
	var category model.Categories
	if err := r.db.
		Model(&category).
		Where("category_id = ? and enable = 1", uuid).
		Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Nếu get để làm việc
func (r *Repo) GetCategoryByName(name string) (*[]model.Categories, error) {
	var category []model.Categories
	if err := r.db.
		First(&category, "category_name = ?", name).
		Error; err != nil {
		return &[]model.Categories{}, err
	}
	return &category, nil
}

// Nếu get để làm việc
func (r *Repo) GetCategoryByEnable(enable int) (*[]model.Categories, error) {
	var categories []model.Categories
	if err := r.db.
		Find(&categories, "enable = ?", enable).
		Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

// Nếu get để làm việc
func (r *Repo) GetCategoryCreatedByUpdateDate(date time.Time) (*[]model.Categories, error) {
	var categories []model.Categories
	if err := r.db.
		Find(&categories, "update_date = ?", date).
		Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

// Get ra để xem
func (r *Repo) GetAllCategory() ([]model.Categories, error) {
	// Hiển thị thông tin tất cả category
	var categories []model.Categories

	if err := r.db.
		Find(&categories).
		Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// Tạo category
func (r *Repo) CreateCategory(category model.Categories) (model.Categories, error) {
	if err := r.db.
		Create(&category).
		Error; err != nil {
		return model.Categories{}, err
	}
	return category, nil
}

// Cập nhật category
func (r *Repo) UpdateCategory(category *model.Categories) (model.Categories, error) {
	// UpdateCategory là hàm cập nhật thông tin category
	if err := r.db.
		Model(&category).
		Where("category_id = ?", category.ID).
		Updates(&category).
		Omit("is_delete").Error; err != nil {
		return model.Categories{}, err
	}
	// Trả về thông tin category đã được cập nhật
	return *category, nil
}

func (r *Repo) EnableCategory(category *model.Categories) error {
	// EnableCategory là hàm hiển thị category
	if err := r.db.
		Where("category_id = ?", category.ID).
		Update("enable", 1).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) DisableCategory(category *model.Categories) error {
	var count int64
	// DisableCategory là hàm ẩn category
	if err := r.db.
		Model(&model.Product{}).
		Where("category_id = ?", &category.ID).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("Không thể xóa danh mục có sản phẩm sử dụng")
	}

	// Nếu không có sản phẩm sử dụng danh mục, thực hiện xóa
	if err := r.db.
		Model(&category).
		Where("category_id = ?", category.ID).
		Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteCategory(category *model.Categories) error {
	// Kiểm tra xem có sản phẩm nào được đặt lịch hẹn thuộc danh mục này không
	var count int64
	if err := r.db.
		Where("category_id = ?", category.ID).Count(&count).Error; err != nil {
		return err
	}

	// Nếu có sản phẩm liên quan, không xóa danh mục mà có thể thực hiện xử lý khác tùy thuộc vào yêu cầu của bạn
	if count > 0 {
		return errors.New("Không thể xóa danh mục có sản phẩm được đặt lịch hẹn./nHãy xóa sản phẩm trước.")
	}

	// Xóa danh mục nếu không có sản phẩm nào được đặt lịch hẹn
	if err := r.db.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
