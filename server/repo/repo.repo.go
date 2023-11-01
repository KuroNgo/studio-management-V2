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
	//  activity_log

	GetActivityLog(log *model.ActivityLog) error
	ViewHistoryUserActitvity(log *[]model.ActivityLog) error
	CreateActivityLog(log *model.ActivityLog) error
	UpdateActivityLog(log *model.ActivityLog) error
	DeleteActivityLog(log *model.ActivityLog) error

	// admin

	GetAdmin(ctx context.Context, admin *model.Admin) error
	UpdateAdmin(ctx context.Context, admin *model.Admin) error

	// order

	GetOrderByID(id int) (model.Order, error)
	GetOrderByOrderDate(date time.Time) ([]model.Order, error)
	GetOrderByOrderDateAndTotalAmount(orderDate time.Time, totalAmount int32) ([]model.Order, error)
	GetOrderByTotalAmount(totalAmount int32) ([]model.Order, error)
	GetAllOrder() ([]model.Order, error)
	GetOrderByUserID(uuid uuid.UUID) ([]model.Order, error)
	GetOrderByEnable(enable int) ([]model.Order, error)
	GetOrderByIsUpdate(isUpdate int) ([]model.Order, error)
	CreateOrder(order *model.Order) error
	UpdateOrder(order *model.Order) error
	DeleteOrder(order *model.Order) error
	RemoveOrder(order *model.Order) error

	//  User

	GetUserProfile(email string) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetUserEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserRole(role string) (*[]model.User, error)
	GetUserAddress(address string) (*[]model.User, error)
	GetUserCreatedAt(createdAt time.Time) (*[]model.User, error)
	GetUserCraetedAtAboutTime(createdAt time.Time) (*[]model.User, error)
	GetUserCreatedAtAboutTime2(lastWeek time.Time, today time.Time) (*[]model.User, error)
	GetUserUpdateUser(lastWeek time.Time, today time.Time) (*[]model.User, error)
	GetUserDeleteUser(deleteUser int) (model.User, error)
	GetUserPhone(phone string) (model.User, error)
	GetUserID(userID uuid.UUID) (*model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUserORInsert(user *model.User) (model.User, error)
	UpdateUser(currentUser *model.User) (*model.User, error)
	DeactivateUser(userID uuid.UUID, currentUser model.User) error
	DeleteUser(user *model.User) error
	FindUserByID(uuid string) (*model.User, error)

	// category

	GetCategoryByID(uuid string) (model.Category, error)
	GetCategoryByID2(uuid string) (*model.Category, error)
	GetCategoryByName(name string) (*[]model.Category, error)
	GetCategoryByEnable(enable int) (*[]model.Category, error)
	GetCategoryCreatedByUpdateDate(date time.Time) (*[]model.Category, error)
	GetAllCategory() ([]model.Category, error)
	CreateCategory(category model.Category) (model.Category, error)
	UpdateCategory(category *model.Category) (model.Category, error)
	EnableCategory(category *model.Category) error
	DisableCategory(category *model.Category) error
	DeleteCategory(category *model.Category) error

	// product

	GetProductsByCategory(categoryID string) ([]model.Product, error)
	GetProductByName(name string) ([]model.Product, error)
	GetProductByPrice(minPrice int32, maxPrice int32) ([]model.Product, error)
	GetProductByWhoUpdate(person string) ([]model.Product, error)
	GetProductByUpdateDate(date time.Time) ([]model.Product, error)
	GetAllProduct() ([]model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	UpdateEnable(enable int) error
	Disable() error
	Enable() error
	DeleteProduct(product *model.Product) error
}
