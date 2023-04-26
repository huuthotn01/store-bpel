package event_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/event_service"
	"store-bpel/event_service/schema"
)

func (c *eventBffController) UploadImage(ctx context.Context, request *event_service.UploadImageRequest) error {
	return c.eventAdapter.UploadImage(ctx, &schema.UploadImageRequest{
		EventId:  request.EventId,
		ImageUrl: request.ImageUrl,
	})
}
