package repo

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"

	"co-studio-e-commerce/model"
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
	GetOrder(ctx context.Context, order *model.Order) error
	GetAllOrder(ctx context.Context, order *[]model.Order) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	DeleteOrder(ctx context.Context, order *model.Order) error
	RemoveOrder(ctx context.Context, order *model.Order) error

	//
	// User
	GetUser(user *model.User) error
	GetAllUser(user model.User) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
	ChangeUserStatus(user *model.User) error
	GetUserID(id int) (model.User, error)
	GetUserEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserRole(role string) (model.User, error)
	GetUserAddress(address string) (model.User, error)
	GetUserCreateUser(create_user string) (model.User, error)
	GetUserUpdateUser(update_user string) (model.User, error)
	GetUserCreateTime(create_time string) (model.User, error)
	GetUserDeleteUser(delete_user int) (model.User, error)
	GetUserPhone(phone string) (model.User, error)

	// category
	GetCategory(ctx context.Context, category *model.Categories) error
	GetAllCategory(ctx context.Context, category *[]model.Categories) error
	CreateCategory(ctx context.Context, category *model.Categories) error
	UpdateCategory(ctx context.Context, category *model.Categories) error
	DeleteCategory(ctx context.Context, category *model.Categories) error
}
