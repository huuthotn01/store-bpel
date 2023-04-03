package repository

import (
	"time"

	"gorm.io/gorm"
)

type goodsServiceRepository struct {
	db                 *gorm.DB
	goodsTableName     string
	goodsImgTableName  string
	goodsInWhTableName string
}

type GoodsModel struct {
	GoodsCode    string
	GoodsSize    string
	GoodsColor   string
	GoodsName    string
	GoodsType    string
	GoodsGender  int
	GoodsAge     string
	Manufacturer string
	IsForSale    int
	UnitPrice    int
	UnitCost     int
	Description  string
}

type GoodsImg struct {
	GoodsCode  string
	GoodsColor string
	GoodsImg   string
	IsDefault  bool
}

type GoodsInWh struct {
	GoodsCode   string
	GoodsSize   string
	GoodsColor  string
	WhCode      string
	Quantity    int
	CreatedDate time.Time
	UpdatedDate time.Time
}
