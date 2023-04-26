package order_service

import (
	"context"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/internal/adapter"
	"store-bpel/bff/shared_bff/schema/order_service"
)

type IOrderBffController interface {
	GetShippingFee(ctx context.Context, request *order_service.Address) (*order_service.GetShipFeeResponseData, error)
	UpdateOnlineOrdersStatus(ctx context.Context, request *order_service.UpdateOnlineOrdersStatusRequest) error
}

type orderBffController struct {
	cfg          *config.Config
	orderAdapter adapter.IOrderServiceAdapter
}

func NewController(cfg *config.Config) IOrderBffController {
	// init order adapter
	orderAdapter := adapter.NewOrderAdapter(cfg)

	return &orderBffController{
		cfg:          cfg,
		orderAdapter: orderAdapter,
	}
}
