package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/goods_service/adapter"
	"store-bpel/goods_service/config"
	repo "store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

type IGoodsServiceController interface {
	GetGoods(ctx context.Context) (*schema.GetGoodsResponse, error)
}

type goodsServiceController struct{
	cfg *config.Config
	repository repo.IGoodsServiceRepository

	warehouseServiceAdapter adapter.IWarehouseServiceAdapter
	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IGoodsServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init warehouse service adapter
	whAdapter := adapter.NewWarehouseAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &goodsServiceController{
		cfg: cfg,
		repository: repository,
		warehouseServiceAdapter: whAdapter,
		kafkaAdapter: kafkaAdapter,
	}
}