package repository

import (
	"gorm.io/gorm"
)

type cartServiceRepository struct {
	db             *gorm.DB
	cartTableName  string
	goodsTableName string
}

type CartModel struct {
	CartId     int
	CustomerId string
}

type GetCartModel struct {
	CartId int
	Goods  []*GoodsModel
}

type GoodsModel struct {
	CartId     int
	GoodsId    string
	GoodsSize  string
	GoodsColor string
	Quantity   int
}
