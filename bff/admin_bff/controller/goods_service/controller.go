package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

type IGoodsBffController interface {
	GetGoods(ctx context.Context) ([]*goods_service.GetGoodsResponseData, error)
	GetGoodsDetail(ctx context.Context, request *goods_service.GetGoodsDetailRequest) (*goods_service.GetGoodsResponseData, error)
	AddGoods(ctx context.Context, request []*goods_service.AddGoodsRequestData) error
	Import(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	Export(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	ReturnManufacturer(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	CustomerReturn(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error
	GetWarehouseByGoods(ctx context.Context, request *goods_service.GetWarehouseByGoodsRequest) ([]*goods_service.GetWarehouseResponseData, error)
	UpdateGoods(ctx context.Context, request []*goods_service.UpdateGoodsRequestData) error
	UploadImage(ctx context.Context, goodsId, goodsColor, url string, isDefault bool) error
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
