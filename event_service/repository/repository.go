package repository

import (
	"context"

	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	GetAllEvent(ctx context.Context) ([]*EventModel, error)
	GetGoods(ctx context.Context, eventId int) ([]string, error)
	GetEvent(ctx context.Context, eventId int) (*EventModel, error)
	AddEvent(ctx context.Context, data *AddEventData) error
	AddGoods(ctx context.Context, eventId int, listGoods []string) error
	UpdateEvent(ctx context.Context, data *UpdateEventData) error
	DeleteGoods(ctx context.Context, eventId int) error
	DeleteEvent(ctx context.Context, eventId int) error
	GetEventByGoods(ctx context.Context, goodsId string) ([]*EventModel, error)
}

func NewRepository(db *gorm.DB) IEventServiceRepository {
	return &eventServiceRepository{
		db:             db,
		eventTableName: "event",
		goodsTableName: "goods",
	}
}

type (
	AddEventData struct {
		EventModel *EventModel
		GoodsList  []string
	}

	UpdateEventData struct {
		EventModel *EventModel
		EventId    int
		GoodsList  []string
	}
)

func (r *eventServiceRepository) GetAllEvent(ctx context.Context) ([]*EventModel, error) {

	var result []*EventModel
	query := r.db.WithContext(ctx).Table(r.eventTableName).Find(&result)

	return result, query.Error
}

func (r *eventServiceRepository) GetGoods(ctx context.Context, eventId int) ([]string, error) {
	var goodsList []*GoodsModel

	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = ?", eventId).Find(&goodsList)

	var result []string
	for _, goods := range goodsList {
		result = append(result, goods.GoodsId)
	}

	return result, query.Error
}

func (r *eventServiceRepository) GetEvent(ctx context.Context, eventId int) (*EventModel, error) {
	var result *EventModel

	query := r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?", eventId).First(&result)

	return result, query.Error
}

func (r *eventServiceRepository) AddEvent(ctx context.Context, data *AddEventData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(r.eventTableName).Create(data.EventModel).Error
		if err != nil {
			return err
		}

		var goodsList []*GoodsModel
		for _, goods := range data.GoodsList {
			goodsList = append(goodsList, &GoodsModel{
				EventId: data.EventModel.EventId,
				GoodsId: goods,
			})
		}
		return tx.Table(r.goodsTableName).Create(goodsList).Error
	})
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

func (r *eventServiceRepository) UpdateEvent(ctx context.Context, data *UpdateEventData) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(r.eventTableName).Where("event_id = ?", data.EventId).Updates(data.EventModel).Error
		if err != nil {
			return err
		}

		err = tx.Table(r.goodsTableName).Where("event_id = ?", data.EventId).Delete(data.EventId).Error
		if err != nil {
			return err
		}

		var goodsList []*GoodsModel
		for _, goods := range data.GoodsList {
			goodsList = append(goodsList, &GoodsModel{
				EventId: data.EventId,
				GoodsId: goods,
			})
		}
		return tx.Table(r.goodsTableName).Create(goodsList).Error
	})
}

func (r *eventServiceRepository) DeleteGoods(ctx context.Context, eventId int) error {
	//delete all current goods
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = ?", eventId).Delete(eventId).Error
}

func (r *eventServiceRepository) DeleteEvent(ctx context.Context, eventId int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(r.eventTableName).Where("event_id = ?", eventId).Delete(eventId).Error
		if err != nil {
			return err
		}

		return tx.Table(r.goodsTableName).Where("event_id = ?", eventId).Delete(eventId).Error
	})
}

func (r *eventServiceRepository) GetEventByGoods(ctx context.Context, goodsId string) ([]*EventModel, error) {
	var eventIdList []string
	err := r.db.WithContext(ctx).Table(r.goodsTableName).Select("event_id").Where("goods_id = ?", goodsId).
		Order("discount desc").Find(&eventIdList).Error
	if err != nil {
		return nil, err
	}

	var result []*EventModel
	for _, eventId := range eventIdList {
		var event *EventModel
		err = r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?", eventId).First(&event).Error
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}

	return result, nil
}
