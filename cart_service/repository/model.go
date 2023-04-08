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
	CartId     string
	CustomerId string
}

type GetCartModel struct {
	CartId string
	Goods  []*GoodsModel
}

type GoodsModel struct {
	CartId     string
	GoodsId    string
	GoodsSize  string
	GoodsColor string
	Quantity   int
}
