package customer_service

import (
	"context"
	"store-bpel/bff/customer_bff/adapter"
	"store-bpel/bff/customer_bff/config"
	"store-bpel/bff/customer_bff/schema/customer_service"
)

type ICustomerBffController interface {
	AddCustomer(ctx context.Context, request *customer_service.AddCustomerRequest) error
	UpdateCustomer(ctx context.Context, request *customer_service.UpdateCustomerInfoRequest) error
	GetCustomer(ctx context.Context, request *customer_service.GetCustomerInfoRequest) (*customer_service.GetCustomerInfoResponseData, error)
}

type customerBffController struct {
	cfg             *config.Config
	customerAdapter adapter.ICustomerServiceAdapter
}

func NewController(cfg *config.Config) ICustomerBffController {
	// init customer adapter
	customerAdapter := adapter.NewCustomerAdapter(cfg)

	return &customerBffController{
		cfg:             cfg,
		customerAdapter: customerAdapter,
	}
}
