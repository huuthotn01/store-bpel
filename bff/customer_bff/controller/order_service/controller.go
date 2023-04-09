package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/adapter"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/order_service"
)

type IOrderBffController interface {
	CreateOnlineOrder(ctx context.Context, request *order_service.MakeOnlineOrderRequest) error
	GetListOrderCustomer(ctx context.Context, request *order_service.GetListOrderCustomerRequest) ([]*order_service.GetListOrderCustomerResponseData, error)
	GetOrderCustomerDetail(ctx context.Context, request *order_service.GetOrderDetailCustomerRequest) (*order_service.GetOrderDetailCustomerResponseData, error)
	GetOnlineOrdersStatus(ctx context.Context, request *order_service.GetOnlineOrdersStatusRequest) ([]*order_service.GetOnlineOrdersStatusResponseData, error)
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
