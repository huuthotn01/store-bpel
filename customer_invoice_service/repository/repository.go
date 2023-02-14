package repository

import (
	"gorm.io/gorm"
)

type ICustomerInvoiceServiceRepository interface {
	
}

func NewRepository(db *gorm.DB) ICustomerInvoiceServiceRepository {
	return &customerInvoiceServiceRepository{
		db: db,
	}
}
