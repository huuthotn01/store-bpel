package repository

import (
	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) IEventServiceRepository {
	return &eventServiceRepository{
		db: db,
	}
}
