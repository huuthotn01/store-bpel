package event_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/event_service"
)

func (c *eventBffController) DeleteEvent(ctx context.Context, request *event_service.DeleteEventRequest) error {

	return c.eventAdapter.DeleteEvent(ctx, request.EventId)

}
