package repository

import (
	"gorm.io/gorm"
)

type IStatisticServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) IStatisticServiceRepository {
	return &statisticServiceRepository{
		db: db,
	}
}
