package repository

import (
	"context"
	"gorm.io/gorm"
)

type IOrderServiceRepository interface {
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

func (r *orderServiceRepository) GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error) {
	var result []*OrderStateModel
	query := r.db.WithContext(ctx).Table(r.orderStateTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) UpdateOrderState(ctx context.Context, orderState *OrderStateModel) error {
	return r.db.WithContext(ctx).Table(r.orderStateTableName).Select("order_code", "state").Create(&orderState).Error
}
