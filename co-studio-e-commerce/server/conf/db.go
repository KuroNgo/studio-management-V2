package conf

import (
	"co-studio-e-commerce/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DbDefault *gorm.DB

// sử dụng singleton pattern để tạo một connection duy nhất đến database
// khi ứng dụng lớn hơn thì không nên sử dụng singleton pattern
// thay vào đó nên sử dụng connection pool
func (a *App) GetDB() *gorm.DB {
	if DbDefault == nil {
		return a.initDB()
	}
	return DbDefault
}

func (a *App) initDB() *gorm.DB {
	//  Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	// Mở kết nối đến PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Categories{},
		&model.Order{},
		&model.OrderDetail{},
		&model.Admin{},
		&model.ActivityLog{},
		&model.Login_Session{},
		&model.PaymentInformation{},
		//&model.Post{},
		&model.Server{},
		&model.ShoppingCart{},
		&model.Reviews{},
	)

	return db
}
