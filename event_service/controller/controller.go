package controller

import (
	"gorm.io/gorm"
	"store-bpel/event_service/adapter"
	"store-bpel/event_service/config"
	repo "store-bpel/event_service/repository"
	// "store-bpel/event_service/schema"
)

type IEventServiceController interface {
}

type eventServiceController struct {
	cfg        *config.Config
	repository repo.IEventServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IEventServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &eventServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
