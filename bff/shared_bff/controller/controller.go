package controller

import (
	"context"
	"store-bpel/bff/shared_bff/adapter"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema"
)

type IGoodsBffController interface {
	GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error)
	GetGoodsDetail(ctx context.Context, request *schema.GetGoodsDetailRequest) (*schema.GetGoodsResponseData, error)
	CreateTransfer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error
	CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error)
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
