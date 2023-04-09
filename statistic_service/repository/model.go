package repository

import (
	"gorm.io/gorm"
)

type statisticServiceRepository struct {
	db              *gorm.DB
	goodsTableName  string
	ordersTableName string
}

type GoodsModel struct {
	GoodsCode   string
	GoodsSize   string
	GoodsColor  string
	GoodsType   string
	GoodsGender int
	GoodsCost   int
	UnitPrice   int
	Quantity    int
	OrderCode   string
}

type OrdersModel struct {
	OrderCode       string
	TransactionDate string
	ShopCode        string
}
