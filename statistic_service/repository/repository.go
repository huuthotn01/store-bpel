package repository

import (
	"context"

	"gorm.io/gorm"
)

type IStatisticServiceRepository interface {
	AddOrderData(ctx context.Context, goodsModel []*GoodsModel, orderModel *OrdersModel) error
	GetOverallStat(ctx context.Context, start, end string, goodsId string, branchId []string, gender []int, goodsType []string) ([]*OverallStatData, error)
}

type OverallStatData struct {
	Revenue int
	Profit  int
	Date    string
}

func NewRepository(db *gorm.DB) IStatisticServiceRepository {
	return &statisticServiceRepository{
		db:              db,
		goodsTableName:  "goods",
		ordersTableName: "orders",
	}
}

func (r *statisticServiceRepository) AddOrderData(ctx context.Context, goodsModel []*GoodsModel, orderModel *OrdersModel) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(r.ordersTableName).Create(&orderModel).Error
		if err != nil {
			return err
		}

		return tx.Table(r.goodsTableName).Create(&goodsModel).Error
	})
}

func (r *statisticServiceRepository) GetOverallStat(ctx context.Context, start, end string, goodsId string, branchId []string, gender []int, goodsType []string) ([]*OverallStatData, error) {
	var result []*OverallStatData
	query := r.db.WithContext(ctx).Table(r.ordersTableName).Joins("left join goods g on orders.order_code = g.order_code").
		Where("orders.transaction_date >= ? and orders.transaction_date <= ?", start, end).
		Group("date(orders.transaction_date)").
		Select("sum(g.quantity * g.unit_price) as revenue, sum(g.quantity * (g.unit_price - g.goods_cost)) as profit, date(orders.transaction_date) as date")
	if goodsId != "" {
		query = query.Where("g.goods_code = ?", goodsId)
	}
	if branchId != nil && len(branchId) > 0 {
		query = query.Where("orders.shop_code in ?", branchId)
	}
	if gender != nil && len(gender) > 0 {
		query = query.Where("g.goods_gender in ?", gender)
	}
	if goodsType != nil && len(goodsType) > 0 {
		query = query.Where("g.goods_type in ?", goodsType)
	}
	return result, query.Find(&result).Error
}
