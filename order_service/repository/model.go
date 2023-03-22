package repository

import (
	"gorm.io/gorm"
)

type orderServiceRepository struct {
	db *gorm.DB
}
