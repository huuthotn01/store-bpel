package event_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/event_service"
)

func (c *eventBffController) DeleteImage(ctx context.Context, request *event_service.DeleteImageRequest) error {
	return c.eventAdapter.DeleteImage(ctx, request.EventId)
}
