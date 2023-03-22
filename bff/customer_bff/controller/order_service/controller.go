package order_service

import (
	"store-bpel/bff/customer_bff/adapter"
	"store-bpel/bff/customer_bff/config"
)

type IOrderBffController interface {
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
