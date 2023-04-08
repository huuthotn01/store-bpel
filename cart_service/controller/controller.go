package controller

import (
	"context"
	"store-bpel/cart_service/adapter"
	"store-bpel/cart_service/config"
	repo "store-bpel/cart_service/repository"
	"store-bpel/cart_service/schema"

	"gorm.io/gorm"
)

type ICartServiceController interface {
	AddCart(ctx context.Context, request string) error
	GetCart(ctx context.Context, request string) (*schema.CartData, error)
	AddGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error
	DeleteGoods(ctx context.Context, cartId string, request []*schema.DeleteGoodsRequest) error
	UpdateGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error
	DeleteAllGoods(ctx context.Context, cartId string) error
}

type cartServiceController struct {
	cfg        *config.Config
	repository repo.ICartServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
	goodsAdapter adapter.IGoodsServiceAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) ICartServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	goodsAdapter := adapter.NewGoodsAdapter(cfg)

	return &cartServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
		goodsAdapter: goodsAdapter,
	}
}
