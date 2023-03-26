package repository

import (
	"context"
	"gorm.io/gorm"
)

type IOrderServiceRepository interface {
	CreateOnlineOrder(ctx context.Context, data *OnlineOrdersData) error
	GetOnlineOrders(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error)
	GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error)
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
		OrderData       *OrdersModel
		OnlineOrderData *OnlineOrdersModel
	}
)

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

func (r *orderServiceRepository) GetOnlineOrders(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error) {
	return nil, nil
}

func (r *orderServiceRepository) GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error) {
	var result []*OrderStateModel
	query := r.db.WithContext(ctx).Table(r.orderStateTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) UpdateOrderState(ctx context.Context, orderState *OrderStateModel) error {
	return r.db.WithContext(ctx).Table(r.orderStateTableName).Select("order_code", "state").Create(&orderState).Error
}
