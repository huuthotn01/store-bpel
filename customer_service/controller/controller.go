package controller

import (
	"gorm.io/gorm"
	"store-bpel/customer_service/adapter"
	"store-bpel/customer_service/config"
	repo "store-bpel/customer_service/repository"
	"store-bpel/customer_service/schema"
)

type ICustomerServiceController interface {
	
}

type customerServiceController struct{
	cfg *config.Config
	repository repo.ICustomerServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICustomerServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &customerServiceController{
		cfg: cfg,
		repository: repository,
		kafkaAdapter: kafkaAdapter,
	}
}
