package event_service

import (
	"context"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/internal/adapter"
	"store-bpel/bff/admin_bff/schema/event_service"
)

type IEventBffController interface {
	AddEvent(ctx context.Context, request *event_service.AddEventRequest) error
	UpdateEvent(ctx context.Context, request *event_service.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, request *event_service.DeleteEventRequest) error
	UploadImage(ctx context.Context, request *event_service.UploadImageRequest) error
	DeleteImage(ctx context.Context, request *event_service.DeleteImageRequest) error
}

type eventBffController struct {
	cfg          *config.Config
	eventAdapter adapter.IEventServiceAdapter
}

func NewController(cfg *config.Config) IEventBffController {
	// init customer adapter
	eventAdapter := adapter.NewEventAdapter(cfg)

	return &eventBffController{
		cfg:          cfg,
		eventAdapter: eventAdapter,
	}
}
