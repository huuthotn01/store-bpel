package event_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/event_service"
	"store-bpel/event_service/schema"
)

func (c *eventBffController) UpdateEvent(ctx context.Context, request *event_service.UpdateEventRequest) error {

	// Call event adapter
	data := &schema.UpdateEventRequest{
		Name:      request.Name,
		Discount:  request.Discount,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Image:     request.Image,
		Goods:     request.Goods,
	}

	return c.eventAdapter.UpdateEvent(ctx, request.EventId, data)

}
