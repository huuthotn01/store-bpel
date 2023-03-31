package event_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/event_service"
)

func (c *eventBffController) GetEventDetail(ctx context.Context, request *event_service.GetEventDetailRequest) (*event_service.GetEventData, error) {
	// Call event adapter
	event, err := c.eventAdapter.GetEventDetail(ctx, request.EventId)
	if err != nil {
		return nil, err
	}

	return &event_service.GetEventData{
		Id:        event.Id,
		Name:      event.Name,
		Discount:  event.Discount,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		Image:     event.Image,
		Goods:     event.Goods,
	}, nil

}
