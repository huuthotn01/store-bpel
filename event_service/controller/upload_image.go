package controller

import (
	"context"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) UploadImage(ctx context.Context, request *schema.UploadImageRequest) error {
	return s.repository.UpdateImage(ctx, request.EventId, request.ImageUrl)
}
