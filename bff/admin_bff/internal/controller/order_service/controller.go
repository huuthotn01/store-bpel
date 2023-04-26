package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/internal/adapter"
	"store-bpel/bff/admin_bff/schema/order_service"
)

type IOrderBffController interface {
	CreateOfflineOrder(ctx context.Context, request *order_service.MakeOfflineOrderRequest) error
	GetOrderDetail(ctx context.Context, request *order_service.GetOrderDetailRequest) (*order_service.GetOrderDetailResponseData, error)
	GetOnlineOrders(ctx context.Context) ([]*order_service.GetOnlineOrdersResponseData, error)
	GetOfflineOrders(ctx context.Context) ([]*order_service.GetOfflineOrdersResponseData, error)
	GetListOrderCustomer(ctx context.Context, request *order_service.GetListOrderCustomerRequest) ([]*order_service.GetListOrderCustomerResponseData, error)
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
