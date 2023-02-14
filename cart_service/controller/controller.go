package controller

import (
	"gorm.io/gorm"
	"store-bpel/cart_service/adapter"
	"store-bpel/cart_service/config"
	repo "store-bpel/cart_service/repository"
)

type ICartServiceController interface {
}

type cartServiceController struct {
	cfg        *config.Config
	repository repo.ICartServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICartServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &cartServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
