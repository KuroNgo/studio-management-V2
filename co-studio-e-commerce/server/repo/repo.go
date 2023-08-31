package repo

import (
	"context"

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
	// user
	GetUser(ctx context.Context, user *model.User) error
	GetAllUser(ctx context.Context, user *[]model.User) error
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error

	// product
	GetProduct(ctx context.Context, product *model.Product) error
	GetAllProduct(ctx context.Context, product *[]model.Product) error
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, product *model.Product) error

	// order
	GetOrder(ctx context.Context, order *model.Order) error
	GetAllOrder(ctx context.Context, order *[]model.Order) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	DeleteOrder(ctx context.Context, order *model.Order) error

	// order_detail
	GetOrderDetail(ctx context.Context, order_detail *model.OrderDetail) error
	GetAllOrderDetail(ctx context.Context, order_detail *[]model.OrderDetail) error
	CreateOrderDetail(ctx context.Context, order_detail *model.OrderDetail) error
	UpdateOrderDetail(ctx context.Context, order_detail *model.OrderDetail) error
	DeleteOrderDetail(ctx context.Context, order_detail *model.OrderDetail) error

	// admin
	GetAdmin(ctx context.Context, admin *model.Admin) error
	GetAllAdmin(ctx context.Context, admin *[]model.Admin) error
	CreateAdmin(ctx context.Context, admin *model.Admin) error
	UpdateAdmin(ctx context.Context, admin *model.Admin) error
	DeleteAdmin(ctx context.Context, admin *model.Admin) error

	// category
	GetCategory(ctx context.Context, category *model.Categories) error
	GetAllCategory(ctx context.Context, category *[]model.Categories) error
	CreateCategory(ctx context.Context, category *model.Categories) error
	UpdateCategory(ctx context.Context, category *model.Categories) error
	DeleteCategory(ctx context.Context, category *model.Categories) error

	// cart
	GetCart(ctx context.Context, cart *model.ShoppingCart) error
	GetAllCart(ctx context.Context, cart *[]model.ShoppingCart) error
	CreateCart(ctx context.Context, cart *model.ShoppingCart) error
	UpdateCart(ctx context.Context, cart *model.ShoppingCart) error
	DeleteCart(ctx context.Context, cart *model.ShoppingCart) error

	// review
	GetReview(ctx context.Context, review *model.Reviews) error
	GetAllReview(ctx context.Context, review *[]model.Reviews) error
	CreateReview(ctx context.Context, review *model.Reviews) error
	UpdateReview(ctx context.Context, review *model.Reviews) error
	DeleteReview(ctx context.Context, review *model.Reviews) error

	// payment
	GetPayment(ctx context.Context, payment *model.PaymentInformation) error
	GetAllPayment(ctx context.Context, payment *[]model.PaymentInformation) error
	CreatePayment(ctx context.Context, payment *model.PaymentInformation) error
	UpdatePayment(ctx context.Context, payment *model.PaymentInformation) error
	DeletePayment(ctx context.Context, payment *model.PaymentInformation) error
}
