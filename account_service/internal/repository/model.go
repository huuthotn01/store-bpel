package repository

import (
	"time"

	"gorm.io/gorm"
)

type accountServiceRepository struct {
	db               *gorm.DB
	accountTableName string
}

type AccountModel struct {
	Username    string
	Password    string
	UserRole    int
	Email       string
	IsActivated int
	CreatedAt   time.Time
}
