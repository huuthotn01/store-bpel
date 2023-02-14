package repository

import (
	"gorm.io/gorm"
)

type customerInvoiceServiceRepository struct {
	db *gorm.DB
}
