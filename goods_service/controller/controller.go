package controller

import (
	"context"
	"store-bpel/goods_service/adapter"
	"store-bpel/goods_service/config"
	repo "store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"

	"gorm.io/gorm"
)

type IGoodsServiceController interface {
	GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error)
	GetGoodsDefault(ctx context.Context, request *schema.GetGoodsDefaultRequest) ([]*schema.GetGoodsDefaultResponseData, error)
	GetProductDetail(ctx context.Context, goodsId string) (*schema.GetGoodsDefaultResponseData, error)
	CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error)
	GetDetailGoods(ctx context.Context, goodsId string) (*schema.GetGoodsResponseData, error)
	AddGoods(ctx context.Context, request []*schema.AddGoodsRequest) error
	UpdateGoods(ctx context.Context, request []*schema.UpdateGoodsRequest, goodsId string) error
	DeleteGoods(ctx context.Context, goodsId string) error
	CreateGoodsTransaction(ctx context.Context, request *schema.CreateGoodsTransactionRequest, transactionType string) error
	GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*schema.GetGoodsInWarehouseResponseData, error)
	UploadGoodsImage(ctx context.Context, request *schema.UploadImageRequest) error
}

type goodsServiceController struct {
	cfg        *config.Config
	repository repo.IGoodsServiceRepository

	warehouseServiceAdapter adapter.IWarehouseServiceAdapter
	eventServiceAdapter     adapter.IEventServiceAdapter
	kafkaAdapter            adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IGoodsServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init warehouse service adapter
	whAdapter := adapter.NewWarehouseAdapter(cfg)

	// init event service adapter
	eventAdapter := adapter.NewEventAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter(cfg)

	return &goodsServiceController{
		cfg:                     cfg,
		repository:              repository,
		warehouseServiceAdapter: whAdapter,
		kafkaAdapter:            kafkaAdapter,
		eventServiceAdapter:     eventAdapter,
	}
}
