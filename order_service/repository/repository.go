package repository

import (
	"gorm.io/gorm"
)

type IOrderServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) IOrderServiceRepository {
	return &orderServiceRepository{
		db: db,
	}
}
