package cart_service

import (
	"context"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/internal/adapter"
	"store-bpel/bff/customer_bff/schema/cart_service"
)

type ICartBffController interface {
	GetCart(ctx context.Context, userId string) (*cart_service.CartData, error)
	AddGoods(ctx context.Context, request *cart_service.AddGoodsRequest) error
	UpdateGoods(ctx context.Context, request *cart_service.UpdateGoodsRequest) error
	DeleteGoods(ctx context.Context, request *cart_service.DeleteGoodsRequest) error
	DeleteAllGoods(ctx context.Context, cartId string) error
}

type cartBffController struct {
	cfg         *config.Config
	cartAdapter adapter.ICartServiceAdapter
}

func NewController(cfg *config.Config) ICartBffController {
	// init customer adapter
	cartAdapter := adapter.NewCartAdapter(cfg)

	return &cartBffController{
		cfg:         cfg,
		cartAdapter: cartAdapter,
	}
}
