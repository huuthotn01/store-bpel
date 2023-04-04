package controller

import (
	"gorm.io/gorm"
	"store-bpel/statistic_service/adapter"
	"store-bpel/statistic_service/config"
	repo "store-bpel/statistic_service/repository"
	"store-bpel/statistic_service/schema"
)

type IStatisticServiceController interface {
	
}

type statisticServiceController struct{
	cfg *config.Config
	repository repo.IStatisticServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IStatisticServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &statisticServiceController{
		cfg: cfg,
		repository: repository,
		kafkaAdapter: kafkaAdapter,
	}
}
