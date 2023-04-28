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
	Otp         string
	OtpTimeout  time.Time
	IsActivated int
	CreatedAt   time.Time
}
