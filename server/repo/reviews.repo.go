package repo

import (
	"co-studio-e-commerce/model"
	"github.com/google/uuid"
	"time"
)

// Lấy tất cả bình luận trong một bài viết(blog)
func (r *Repo) GetAllReviewInOneProduct(id uuid.UUID) ([]model.Reviews, error) {
	var review []model.Reviews
	if err := r.db.Where("product_id = ?", id.String()).Find(review).Error; err != nil {
		return []model.Reviews{}, err
	}
	return review, nil
}

// Lấy thông tin review bằng ID
func (r *Repo) GetReviewByID(id int32) (model.Reviews, error) {
	var review model.Reviews
	if err := r.db.Model(&review).Where("review_id = ?", id).First(review).Error; err != nil {
		return model.Reviews{}, err
	}
	return review, nil
}

// Lấy thông tin review theo mã sản phẩm
func (r *Repo) GetReviewByProductID(id uuid.UUID) (model.Reviews, error) {
	var review model.Reviews
	if err := r.db.Model(&review).Where("product_id = ?", id.String()).First(review).Error; err != nil {
		return model.Reviews{}, err
	}
	return review, nil
}

func (r *Repo) GetReviewByRating(rate int) ([]model.Reviews, error) {
	var review []model.Reviews
	if err := r.db.Model(&review).Where("rating = ?", rate).Find(review).Error; err != nil {
		return []model.Reviews{}, err
	}
	return review, nil
}

func (r *Repo) GetReviewByReviewDate(date time.Time) ([]model.Reviews, error) {
	var review []model.Reviews
	if err := r.db.Model(&review).Where("review_date = ?", date).Find(review).Error; err != nil {
		return []model.Reviews{}, err
	}
	return review, nil
}

func (r *Repo) GetReviewByEnable(enable int) ([]model.Reviews, error) {
	var review []model.Reviews
	if err := r.db.Where("enable = ?", enable).Find(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}
