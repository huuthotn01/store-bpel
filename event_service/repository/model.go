package repository

import (
	"gorm.io/gorm"
)

type eventServiceRepository struct {
	db *gorm.DB
}
