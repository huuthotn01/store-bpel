package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/adapter"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/order_service"
)

type IOrderBffController interface {
	CreateOnlineOrder(ctx context.Context, request *order_service.MakeOnlineOrderRequest) error
	GetShippingFee(ctx context.Context, request *order_service.Address) (int, error)
	GetOnlineOrdersStatus(ctx context.Context, request *order_service.GetOnlineOrdersStatusRequest) ([]*order_service.GetOnlineOrdersStatusResponseData, error)
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
