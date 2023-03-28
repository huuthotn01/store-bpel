package repository

import (
	"context"
	"strconv"

	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	GetAllEvent(ctx context.Context) ([]*EventModel, error)
	GetGoods(ctx context.Context, eventId int) ([]string, error)
	GetEvent(ctx context.Context, eventId int) (*EventModel, error)
	AddEvent(ctx context.Context, data *EventModel) (int, error)
	AddGoods(ctx context.Context, eventId int, listGoods []string) error
	UpdateEvent(ctx context.Context, eventId int, data *EventModel) error
	DeleteGoods(ctx context.Context, eventId int) error
	DeleteEvent(ctx context.Context, eventId int) error
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

func (r *eventServiceRepository) GetGoods(ctx context.Context, eventId int) ([]string, error) {
	var goodsList []*GoodsModel

	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = '" + strconv.Itoa(eventId) + "'").Find(&goodsList)

	var result []string
	for _, goods := range goodsList {
		result = append(result, goods.GoodsId)
	}

	return result, query.Error
}

func (r *eventServiceRepository) GetEvent(ctx context.Context, eventId int) (*EventModel, error) {
	var result *EventModel

	query := r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = '" + strconv.Itoa(eventId) + "'").First(&result)

	return result, query.Error
}

func (r *eventServiceRepository) AddEvent(ctx context.Context, data *EventModel) (int, error) {
	result := r.db.WithContext(ctx).Table(r.eventTableName).Create(data)
	return data.EventId, result.Error
}

func (r *eventServiceRepository) AddGoods(ctx context.Context, eventId int, listGoods []string) error {
	var data []*GoodsModel
	for _, goods := range listGoods {
		data = append(data, &GoodsModel{
			EventId: eventId,
			GoodsId: goods,
		})
	}
	return r.db.WithContext(ctx).Table(r.goodsTableName).Create(data).Error
}

func (r *eventServiceRepository) UpdateEvent(ctx context.Context, eventId int, data *EventModel) error {
	return r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?", eventId).Updates(data).Error
}

func (r *eventServiceRepository) DeleteGoods(ctx context.Context, eventId int) error {
	//delete all current goods
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = ?", eventId).Delete(eventId).Error
}

func (r *eventServiceRepository) DeleteEvent(ctx context.Context, eventId int) error {
	return r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?", eventId).Delete(eventId).Error
}
