package controller

import (
	"context"
	"store-bpel/customer_service/config"
	repo "store-bpel/customer_service/internal/repository"
	"store-bpel/customer_service/schema"

	"gorm.io/gorm"
)

type ICustomerServiceController interface {
	GetCustomerInfo(ctx context.Context, customerId string) (*schema.GetCustomerInfoResponseData, error)
	UpdateCustomerInfo(ctx context.Context, customerId string, request *schema.UpdateCustomerInfoRequest) error
	AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error
	UploadImage(ctx context.Context, request *schema.UploadImageRequest) error
	DeleteImage(ctx context.Context, username string) error
}

type customerServiceController struct {
	cfg        *config.Config
	repository repo.ICustomerServiceRepository
}

func NewController(cfg *config.Config, db *gorm.DB) ICustomerServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &customerServiceController{
		cfg:        cfg,
		repository: repository,
	}
}
