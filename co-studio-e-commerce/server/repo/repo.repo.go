package repo

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"

	"co-studio-e-commerce/model"
	"github.com/google/uuid"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

type IRepo interface {
	// activity_log
	GetActivityLog(ctx context.Context, log *model.ActivityLog) error
	ViewHistoryUserActitvity(ctx context.Context, log *[]model.ActivityLog) error
	CreateActivityLog(ctx context.Context, log *model.ActivityLog) error
	UpdateActivityLog(ctx context.Context, log *model.ActivityLog) error
	DeleteActivityLog(ctx context.Context, log *model.ActivityLog) error

	// admin
	GetAdmin(ctx context.Context, admin *model.Admin) error
	UpdateAdmin(ctx context.Context, admin *model.Admin) error

	// order
	GetOrderByID(id int) (model.Order, error)
	GetOrderByOrderDate(date time.Time) ([]model.Order, error)
	GetOrderByOrderDateAndTotalAmount(order_date time.Time, total_amount int32) ([]model.Order, error)
	GetOrderByTotalAmount(total_amount int32) ([]model.Order, error)
	GetAllOrder() ([]model.Order, error)
	GetOrderByUserID(uuid uuid.UUID) ([]model.Order, error)
	GetOrderByEnable(enable int) ([]model.Order, error)
	GetOrderByIsUpdate(isUpdate int) ([]model.Order, error)
	CreateOrder(order *model.Order) error
	UpdateOrder(order *model.Order) error
	DeleteOrder(order *model.Order) error
	RemoveOrder(order *model.Order) error

	// User
	GetAllUser() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUserORInsert(user *model.User) (model.User, error)
	UpdateUser(currentUser model.User, user *model.User) (model.User, error)
	DeactivateUser(userID uuid.UUID, currentUser model.User) error
	DeleteUser(user *model.User) error
	GetUserID(userID uuid.UUID) (model.User, error)
	GetUserEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserRole(role string) ([]model.User, error)
	GetUserAddress(address string) ([]model.User, error)
	FindUserByID(uuid string) (*model.User, error)
	GetUserCreateUser(create_user string) ([]model.User, error)
	GetUserUpdateUser(update_user string) (model.User, error)
	GetUserCreateTime(create_time string) (model.User, error)
	GetUserDeleteUser(delete_user int) (model.User, error)
	GetUserPhone(phone string) (model.User, error)

	// category
	GetCategoryByID(uuid uuid.UUID) (model.Categories, error)
	GetCategoryByName(name string) (model.Categories, error)
	GetCategoryByEnable(enable int) ([]model.Categories, error)
	GetCategoryCreatedByUpdateDate(date time.Time) ([]model.Categories, error)
	GetCategory(category model.Categories) error
	GetAllCategory() ([]model.Categories, error)
	CreateCategory(category model.Categories) (model.Categories, error)
	UpdateCategory(category *model.Categories) (model.Categories, error)
	EnableCategory(category *model.Categories) error
	DisableCategory(uuid uuid.UUID, category *model.Categories) error
	DeleteCategory(category *model.Categories) error
}
