package controller

import (
	"context"
	"store-bpel/event_service/config"
	repo "store-bpel/event_service/internal/repository"
	"store-bpel/event_service/schema"

	"gorm.io/gorm"
)

type IEventServiceController interface {
	GetEvent(ctx context.Context) ([]*schema.GetEventData, error)
	GetEventCurrent(ctx context.Context, date int) ([]*schema.GetEventData, error)
	GetEventDetail(ctx context.Context, eventId string) (*schema.GetEventData, error)
	AddEvent(ctx context.Context, data *schema.AddEventRequest) error
	UpdateEvent(ctx context.Context, eventId string, data *schema.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, eventId string) error
	GetEventByGoods(ctx context.Context, goodsId string) ([]*schema.GetEventByGoodsData, error)
	UploadImage(ctx context.Context, request *schema.UploadImageRequest) error
	DeleteImage(ctx context.Context, eventId string) error
}

type eventServiceController struct {
	cfg        *config.Config
	repository repo.IEventServiceRepository
}

func NewController(cfg *config.Config, db *gorm.DB) IEventServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &eventServiceController{
		cfg:        cfg,
		repository: repository,
	}
}
