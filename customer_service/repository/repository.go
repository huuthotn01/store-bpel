package repository

import (
	"gorm.io/gorm"
)

type ICustomerServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) ICustomerServiceRepository {
	return &customerServiceRepository{
		db: db,
	}
}
