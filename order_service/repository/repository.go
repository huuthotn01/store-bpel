package repository

import (
	"context"
	"gorm.io/gorm"
)

type IOrderServiceRepository interface {
	GetOnlineOrders(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error)
	GetOrderByOrderId(ctx context.Context, orderId int) (*OrdersModel, error)
	GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*GoodsModel, error)
	GetOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersModel, error)
	GetOrderDetail(ctx context.Context, privateOrderId int) (*OnlineOrdersResponse, error)
	GetPrivateOrderCode(ctx context.Context, orderId string) (int, error)
	GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error)

	CreateOnlineOrder(ctx context.Context, data *OnlineOrdersData) error
	UpdateOrderState(ctx context.Context, orderState *OrderStateModel) error
}

func NewRepository(db *gorm.DB) IOrderServiceRepository {
	return &orderServiceRepository{
		db:                    db,
		goodsTableName:        "goods",
		ordersTableName:       "orders",
		onlineOrdersTableName: "online_orders",
		storeOrdersTableName:  "store_orders",
		orderStateTableName:   "order_state",
	}
}

type (
	OnlineOrdersData struct {
		PublicOrderCode string
		TransactionDate string
		TotalPrice      int
		OnlineOrder     *OnlineOrdersModel
		Goods           []*GoodsModel
	}

	OnlineOrdersResponse struct {
		OrderGoods      []*GoodsModel
		OrderData       *OrdersModel
		OnlineOrderData *OnlineOrdersModel
		ShippingState   []*OrderStateModel
	}
)

/*
	***************
	| GET METHODS |
	***************
*/
func (r *orderServiceRepository) GetOnlineOrders(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error) {
	// get orders made by customer
	orders, err := r.GetOrdersByCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}

	resp := make([]*OnlineOrdersResponse, 0, len(orders))

	for _, order := range orders {
		detail, err := r.GetOrderDetail(ctx, order.OrderCode)
		if err != nil {
			return nil, err
		}

		resp = append(resp, detail)
	}

	return resp, nil
}

func (r *orderServiceRepository) GetOrderDetail(ctx context.Context, privateOrderId int) (*OnlineOrdersResponse, error) {
	// get order data
	order, err := r.GetOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get online order data
	onlineOrder, err := r.GetOnlineOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get order goods
	goods, err := r.GetOrderGoodsByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get order state
	state, err := r.GetOrderState(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	return &OnlineOrdersResponse{
		OrderData:       order,
		OrderGoods:      goods,
		ShippingState:   state,
		OnlineOrderData: onlineOrder,
	}, nil
}

func (r *orderServiceRepository) GetOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersModel, error) {
	var result []*OnlineOrdersModel
	query := r.db.WithContext(ctx).Table(r.onlineOrdersTableName).Where("customer_id = ?", customerId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetPrivateOrderCode(ctx context.Context, orderId string) (int, error) {
	var result int
	query := r.db.WithContext(ctx).Table(r.ordersTableName).Where("public_order_code = ?", orderId).Select("order_code")
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetOrderByOrderId(ctx context.Context, orderId int) (*OrdersModel, error) {
	var result *OrdersModel
	query := r.db.WithContext(ctx).Table(r.ordersTableName).Where("order_code = ?", orderId)
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOnlineOrderByOrderId(ctx context.Context, orderId int) (*OnlineOrdersModel, error) {
	var result *OnlineOrdersModel
	query := r.db.WithContext(ctx).Table(r.onlineOrdersTableName).Where("order_code = ?", orderId)
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error) {
	var result []*OrderStateModel
	query := r.db.WithContext(ctx).Table(r.orderStateTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

/*
	***************
	| UPDATE METHODS |
	***************
*/
func (r *orderServiceRepository) CreateOnlineOrder(ctx context.Context, data *OnlineOrdersData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// add data to orders table
		orderModel := &OrdersModel{
			TransactionDate: data.TransactionDate,
			TotalPrice:      data.TotalPrice,
			PublicOrderCode: data.PublicOrderCode,
		}
		err := tx.Table(r.ordersTableName).Create(orderModel).Error
		if err != nil {
			return err
		}

		// add data to goods table
		for _, g := range data.Goods {
			g.OrderCode = orderModel.OrderCode
		}
		err = tx.Table(r.goodsTableName).Create(data.Goods).Error
		if err != nil {
			return err
		}

		// add data to online_orders table
		data.OnlineOrder.OrderCode = orderModel.OrderCode
		err = tx.Table(r.onlineOrdersTableName).Create(data.OnlineOrder).Error
		return err
	})
}

func (r *orderServiceRepository) UpdateOrderState(ctx context.Context, orderState *OrderStateModel) error {
	return r.db.WithContext(ctx).Table(r.orderStateTableName).Select("order_code", "state").Create(&orderState).Error
}
