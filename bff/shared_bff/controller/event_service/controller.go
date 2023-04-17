package event_service

import (
	"context"
	"store-bpel/bff/shared_bff/adapter"
	"store-bpel/bff/shared_bff/config"
	"store-bpel/bff/shared_bff/schema/event_service"
)

type IEventBffController interface {
	GetEventDetail(ctx context.Context, request *event_service.GetEventDetailRequest) (*event_service.GetEventData, error)
	GetEvent(ctx context.Context) ([]*event_service.GetEventData, error)
	GetEventCurrent(ctx context.Context, request *event_service.GetEventCurrentRequest) ([]*event_service.GetEventData, error)
	GetEventByGoods(ctx context.Context, request *event_service.GetEventByGoodsRequest) ([]*event_service.GetEventByGoodsData, error)
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
