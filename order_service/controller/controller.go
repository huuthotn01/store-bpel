package controller

import (
	"gorm.io/gorm"
	"store-bpel/order_service/adapter"
	"store-bpel/order_service/config"
	repo "store-bpel/order_service/repository"
)

type IOrderServiceController interface {
}

type orderServiceController struct {
	cfg        *config.Config
	repository repo.IOrderServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IOrderServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &orderServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
