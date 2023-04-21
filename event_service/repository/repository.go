package repository

import (
	"context"
	"sort"

	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	GetAllEvent(ctx context.Context) ([]*EventModel, error)
	GetAllEventCurrent(ctx context.Context, date int) ([]*EventModel, error)
	GetGoods(ctx context.Context, eventId string) ([]string, error)
	GetEvent(ctx context.Context, eventId string) (*EventModel, error)
	AddEvent(ctx context.Context, data *AddEventData) error
	AddGoods(ctx context.Context, eventId string, listGoods []string) error
	UpdateEvent(ctx context.Context, data *UpdateEventData) error
	DeleteGoods(ctx context.Context, eventId string) error
	DeleteEvent(ctx context.Context, eventId string) error
	GetEventByGoods(ctx context.Context, goodsId string) ([]*EventModel, error)
	UpdateImage(ctx context.Context, eventId string, imageUrl string) error
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
		EventId    string
		GoodsList  []string
	}
)

type ByDiscount []*EventModel

func (b ByDiscount) Len() int           { return len(b) }
func (b ByDiscount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByDiscount) Less(i, j int) bool { return b[i].Discount > b[j].Discount }

func (r *eventServiceRepository) UpdateImage(ctx context.Context, eventId string, imageUrl string) error {
	return r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?", eventId).Update("image", imageUrl).Error
}

func (r *eventServiceRepository) GetAllEvent(ctx context.Context) ([]*EventModel, error) {

	var result []*EventModel
	query := r.db.WithContext(ctx).Table(r.eventTableName).Order("created_at DESC").Find(&result)

	return result, query.Error
}

func (r *eventServiceRepository) GetAllEventCurrent(ctx context.Context, date int) ([]*EventModel, error) {

	var result []*EventModel
	query := r.db.WithContext(ctx).Table(r.eventTableName).Where("start_time < NOW() + INTERVAL ? DAY AND end_time > NOW()", date).Find(&result)

	return result, query.Error
}

func (r *eventServiceRepository) GetGoods(ctx context.Context, eventId string) ([]string, error) {
	var goodsList []*GoodsModel

	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = ?", eventId).Find(&goodsList)

	var result []string
	for _, goods := range goodsList {
		result = append(result, goods.GoodsId)
	}

	return result, query.Error
}

func (r *eventServiceRepository) GetEvent(ctx context.Context, eventId string) (*EventModel, error) {
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

func (r *eventServiceRepository) AddGoods(ctx context.Context, eventId string, listGoods []string) error {
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

func (r *eventServiceRepository) DeleteGoods(ctx context.Context, eventId string) error {
	//delete all current goods
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("event_id = ?", eventId).Delete(eventId).Error
}

func (r *eventServiceRepository) DeleteEvent(ctx context.Context, eventId string) error {
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
	err := r.db.WithContext(ctx).Table(r.goodsTableName).Select("event_id").Where("goods_id = ?", goodsId).Find(&eventIdList).Error
	if err != nil {
		return nil, err
	}

	var result []*EventModel
	for _, eventId := range eventIdList {
		var event *EventModel
		err = r.db.WithContext(ctx).Table(r.eventTableName).Where("event_id = ?  AND start_time <= NOW() AND NOW() <= end_time", eventId).First(&event).Error
		if err != nil {
			continue
		}
		result = append(result, event)
	}
	if len(result) != 0 {
		sort.Sort(ByDiscount(result))
		return result[:1], nil
	}
	return result, nil

}
