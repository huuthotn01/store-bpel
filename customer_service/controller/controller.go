package controller

import (
	"context"
	"store-bpel/customer_service/adapter"
	"store-bpel/customer_service/config"
	repo "store-bpel/customer_service/repository"
	"store-bpel/customer_service/schema"

	"gorm.io/gorm"
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
	cartAdapter    adapter.ICartServiceAdapter
	kafkaAdapter   adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICustomerServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init account adapter
	accountAdapter := adapter.NewAccountAdapter(cfg)

	// init cart adapter
	cartAdapter := adapter.NewCartAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &customerServiceController{
		cfg:            cfg,
		repository:     repository,
		accountAdapter: accountAdapter,
		cartAdapter:    cartAdapter,
		kafkaAdapter:   kafkaAdapter,
	}
}
