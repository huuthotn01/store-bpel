package repository

import (
	"gorm.io/gorm"
)

type eventServiceRepository struct {
	db             *gorm.DB
	eventTableName string
	goodsTableName string
}

type EventModel struct {
	EventId   int `gorm:"primarykey"`
	Name      string
	Discount  float32
	StartTime string
	EndTime   string
	Image     string
}

type GoodsModel struct {
	EventId int
	GoodsId string
}
