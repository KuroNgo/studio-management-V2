package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

type IMigration interface {
	Migrate(c *gin.Context)
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{db: db}
}
