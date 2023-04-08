package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/order_service/adapter"
	"store-bpel/order_service/config"
	repo "store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

type IOrderServiceController interface {
	CreateOnlineOrder(ctx context.Context, request *schema.MakeOnlineOrderRequest) error
	GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error)
	GetOrderDetail(ctx context.Context, orderId string) (*schema.GetOrderDetailCustomerResponseData, error)
	GetShipFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error)
	GetOnlineOrdersStatus(ctx context.Context, orderId int) ([]*schema.GetOnlineOrdersStatusResponseData, error)
	UpdateOrderState(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error
}

type orderServiceController struct {
	cfg        *config.Config
	repository repo.IOrderServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IOrderServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &orderServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
