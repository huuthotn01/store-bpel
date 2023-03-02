package repository

import (
	"gorm.io/gorm"
	"time"
)

type accountServiceRepository struct {
	db               *gorm.DB
	accountTableName string
}

type AccountModel struct {
	Username    string
	Password    string
	UserRole    int
	IsActivated int
	CreatedAt   time.Time
}
