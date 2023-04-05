package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/adapter"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/goods_service"
)

type IGoodsBffController interface {
	GetGoods(ctx context.Context) ([]*goods_service.GetGoodsResponseData, error)
	GetGoodsDefault(ctx context.Context, request *goods_service.GetGoodsDefaultRequest) ([]*goods_service.GetGoodsDefaultResponseData, error)
	GetProductsDetail(ctx context.Context, request *goods_service.GetProductsDetailRequest) (*goods_service.GetGoodsDefaultResponseData, error)
	GetGoodsDetail(ctx context.Context, request *goods_service.GetGoodsDetailRequest) ([]*goods_service.GetGoodsResponseData, error)
	CreateTransfer(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	CheckWarehouse(ctx context.Context, request *goods_service.CheckWarehouseRequest) (*goods_service.CheckWarehouseResponseData, error)
}

type goodsBffController struct {
	cfg          *config.Config
	goodsAdapter adapter.IGoodsServiceAdapter
}

func NewController(cfg *config.Config) IGoodsBffController {
	// init customer adapter
	goodsAdapter := adapter.NewGoodsAdapter(cfg)

	return &goodsBffController{
		cfg:          cfg,
		goodsAdapter: goodsAdapter,
	}
}
