package repository

import (
	"context"

	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	GetAllEvent(ctx context.Context) ([]*EventModel, error)
	GetGoods(ctx context.Context, eventId string) ([]*string, error)
	// GetEvent(ctx context.Context, eventId string) (*EventModel, error)
	// AddEvent(ctx context.Context, data *EventModel) error
	// UpdateEvent(ctx context.Context, data *EventModel) error
	// DeleteEvent(ctx context.Context, eventId string) error
}

func NewRepository(db *gorm.DB) IEventServiceRepository {
	return &eventServiceRepository{
		db:             db,
		eventTableName: "event",
		goodsTableName: "goods",
	}
}

func (r *eventServiceRepository) GetAllEvent(ctx context.Context) ([]*EventModel, error) {

	var result []*EventModel
	query := r.db.WithContext(ctx).Table(r.eventTableName).Find(&result)

	return result, query.Error
}

func (r *eventServiceRepository) GetGoods(ctx context.Context, eventId string) ([]*string, error) {

	var goodsList []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = " + eventId).Find(&goodsList)

	var result []*string
	for _, goods := range goodsList {
		result = append(result, &goods.GoodsId)
	}

	return result, query.Error
}
