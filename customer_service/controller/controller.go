package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/customer_service/adapter"
	"store-bpel/customer_service/config"
	repo "store-bpel/customer_service/repository"
	"store-bpel/customer_service/schema"
)

type ICustomerServiceController interface {
	GetCustomerInfo(ctx context.Context, customerId string) (*schema.GetCustomerInfoResponseData, error)
	UpdateCustomerInfo(ctx context.Context, customerId string, request *schema.UpdateCustomerInfoRequest) error
	AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error
}

type customerServiceController struct {
	cfg        *config.Config
	repository repo.ICustomerServiceRepository

	accountAdapter adapter.IAccountServiceAdapter
	kafkaAdapter   adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICustomerServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init account adapter
	accountAdapter := adapter.NewAccountAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &customerServiceController{
		cfg:            cfg,
		repository:     repository,
		accountAdapter: accountAdapter,
		kafkaAdapter:   kafkaAdapter,
	}
}
