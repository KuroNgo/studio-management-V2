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
	GetOrder(ctx context.Context, order *model.Order) error
	GetAllOrder(ctx context.Context, order *[]model.Order) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	DeleteOrder(ctx context.Context, order *model.Order) error
	RemoveOrder(ctx context.Context, order *model.Order) error

	// User
	GetUser(user model.User) error
	GetAllUser() ([]model.User, error)              // used
	CreateUser(user model.User) (model.User, error) // used
	UpdateUserORInsert(user *model.User) (model.User, error)
	UpdateUser(currentUser model.User, user *model.User) (model.User, error) // used
	DeactivateUser(userID uuid.UUID, currentUser model.User) error           // unused
	GetUserID(userID uuid.UUID) (model.User, error)                          // unused
	GetUserEmail(email string) (model.User, error)                           // used
	GetUserByUsername(username string) (model.User, error)                   // used
	GetUserRole(role string) (model.User, error)
	GetUserAddress(address string) (model.User, error)
	GetUserCreateUser(create_user string) (model.User, error)
	GetUserUpdateUser(update_user string) (model.User, error)
	GetUserCreateTime(create_time string) (model.User, error)
	GetUserDeleteUser(delete_user int) (model.User, error)
	GetUserPhone(phone string) (model.User, error)

	// category
	GetCategoryByID(uuid uuid.UUID) (model.Categories, error)
	GetCategoryByName(name string) (model.Categories, error)
	GetCategoryByEnable(enable int) ([]model.Categories, error)
	GetCategoryCreatedByUpdateDate(date time.Time) ([]model.Categories, error)
	GetCategory(category *model.Categories) error
	GetAllCategory() ([]model.Categories, error)
	CreateCategory(category model.Categories) (model.Categories, error)
	UpdateCategory(category *model.Categories) (model.Categories, error)
	EnableCategory(category *model.Categories) error
	DisableCategory(category *model.Categories) error
	DeleteCategory(category *model.Categories) error
}
