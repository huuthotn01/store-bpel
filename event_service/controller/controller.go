package controller

import (
	"context"
	"store-bpel/event_service/adapter"
	"store-bpel/event_service/config"
	repo "store-bpel/event_service/repository"
	"store-bpel/event_service/schema"

	"gorm.io/gorm"
)

type IEventServiceController interface {
	GetEvent(ctx context.Context) ([]*schema.GetEventData, error)
	GetEventCurrent(ctx context.Context, date int) ([]*schema.GetEventData, error)
	GetEventDetail(ctx context.Context, eventId int) (*schema.GetEventData, error)
	AddEvent(ctx context.Context, data *schema.AddEventRequest) error
	UpdateEvent(ctx context.Context, eventId int, data *schema.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, eventId int) error
	GetEventByGoods(ctx context.Context, goodsId string) ([]*schema.GetEventByGoodsData, error)
	UploadImage(ctx context.Context, request *schema.UploadImageRequest) error
	DeleteImage(ctx context.Context, eventId string) error
}

type eventServiceController struct {
	cfg        *config.Config
	repository repo.IEventServiceRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) IEventServiceController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &eventServiceController{
		cfg:          cfg,
		repository:   repository,
		kafkaAdapter: kafkaAdapter,
	}
}
