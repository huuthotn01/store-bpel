package repository

import (
	"gorm.io/gorm"
)

type ICartServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) ICartServiceRepository {
	return &cartServiceRepository{
		db: db,
	}
}
