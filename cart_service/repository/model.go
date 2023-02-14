package repository

import (
	"gorm.io/gorm"
)

type cartServiceRepository struct {
	db *gorm.DB
}
