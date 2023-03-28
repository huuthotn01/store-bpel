package repository

import (
	"context"

	"gorm.io/gorm"
)

type IEventServiceRepository interface {
	GetAllEvent(ctx context.Context) ([]*EventModel, error)
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
	queryContext := r.db.WithContext(ctx).Table(r.eventTableName)

	return result, queryContext.Find(&result).Error
}
