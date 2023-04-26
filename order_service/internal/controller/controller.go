package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/library/kafka_lib"
	"store-bpel/order_service/config"
	"store-bpel/order_service/internal/adapter"
	repo "store-bpel/order_service/internal/repository"
	"store-bpel/order_service/schema"
)

type IOrderServiceController interface {
	// Customer
	CreateOnlineOrder(ctx context.Context, request *schema.MakeOnlineOrderRequest) error
	GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error)
	GetOrderDetail(ctx context.Context, orderId string) (*schema.GetOrderDetailCustomerResponseData, error)
	GetOnlineOrdersStatus(ctx context.Context, orderId string) ([]*schema.GetOnlineOrdersStatusResponseData, error)

	// Admin
	CreateOfflineOrder(ctx context.Context, request *schema.MakeOfflineOrderRequest) error
	GetOfflineOrders(ctx context.Context) ([]*schema.GetOfflineOrdersResponseData, error)
	GetOnlineOrders(ctx context.Context) ([]*schema.GetOnlineOrdersResponseData, error)
	GetOrderDetailAdmin(ctx context.Context, orderId int) (*schema.GetOrderDetailAdminResponseData, error)

	// Shared
	GetShipFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error)
	UpdateOrderState(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error
	GetBestSellingGoods(ctx context.Context) ([]string, error)
}

type orderServiceController struct {
	cfg        *config.Config
	repository repo.IOrderServiceRepository

	goodsAdapter adapter.IGoodsServiceAdapter
	kafkaAdapter kafka_lib.IKafkaLib
}

func NewController(cfg *config.Config, db *gorm.DB) IOrderServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init goods adapter
	goodsAdapter := adapter.NewGoodsAdapter(cfg)

	// init kafka adapter
	kafkaAdapter := kafka_lib.NewKafkaLib()

	return &orderServiceController{
		cfg:          cfg,
		repository:   repository,
		goodsAdapter: goodsAdapter,
		kafkaAdapter: kafkaAdapter,
	}
}
