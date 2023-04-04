package repository

import (
	"gorm.io/gorm"
)

type statisticServiceRepository struct {
	db *gorm.DB
}
