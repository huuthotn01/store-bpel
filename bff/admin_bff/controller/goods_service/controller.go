package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

type IGoodsBffController interface {
	AddGoods(ctx context.Context, request *goods_service.AddGoodsRequest) error
	Import(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	Export(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	ReturnManufacturer(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	CustomerReturn(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
}

type goodsBffController struct {
	cfg          *config.Config
	goodsAdapter adapter.IGoodsServiceAdapter
}

func NewController(cfg *config.Config) IGoodsBffController {
	// init branch adapter
	goodsAdapter := adapter.NewGoodsAdapter(cfg)

	return &goodsBffController{
		cfg:          cfg,
		goodsAdapter: goodsAdapter,
	}
}
