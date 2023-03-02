package repository

import (
	"gorm.io/gorm"
)

type customerServiceRepository struct {
	db *gorm.DB
}
