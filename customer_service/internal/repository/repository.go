package repository

import (
	"context"
	"gorm.io/gorm"
)

type ICustomerServiceRepository interface {
	GetCustomerInfo(ctx context.Context, customerId string) (*CustomerModel, error)
	UpdateCustomerInfo(ctx context.Context, data *CustomerModel) error
	AddCustomer(ctx context.Context, data *CustomerModel) error
	UpdateCustomerImage(ctx context.Context, customerId, imageUrl string) error
}

func NewRepository(db *gorm.DB) ICustomerServiceRepository {
	return &customerServiceRepository{
		db:                db,
		customerTableName: "customer",
	}
}

func (r *customerServiceRepository) UpdateCustomerImage(ctx context.Context, customerId, imageUrl string) error {
	return r.db.WithContext(ctx).Table(r.customerTableName).Where("username = ?", customerId).Update("image", imageUrl).Error
}

func (r *customerServiceRepository) GetCustomerInfo(ctx context.Context, customerId string) (*CustomerModel, error) {
	var result *CustomerModel
	query := r.db.WithContext(ctx).Table(r.customerTableName).Where(&CustomerModel{Username: customerId})
	return result, query.First(&result).Error
}

func (r *customerServiceRepository) UpdateCustomerInfo(ctx context.Context, data *CustomerModel) error {
	return r.db.WithContext(ctx).Table(r.customerTableName).Where(&CustomerModel{Username: data.Username}).Updates(&data).Error
}

func (r *customerServiceRepository) AddCustomer(ctx context.Context, data *CustomerModel) error {
	return r.db.WithContext(ctx).Table(r.customerTableName).Create(&data).Error
}
