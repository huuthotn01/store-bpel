package repository

import (
	"time"

	"gorm.io/gorm"
)

type eventServiceRepository struct {
	db             *gorm.DB
	eventTableName string
	goodsTableName string
}

type EventModel struct {
	EventId   string
	Name      string
	Discount  float32
	StartTime time.Time
	EndTime   time.Time
	Image     string
}

type GoodsModel struct {
	EventId string
	GoodsId string
}
