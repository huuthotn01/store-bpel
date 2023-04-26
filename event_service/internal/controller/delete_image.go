package controller

import "context"

func (s *eventServiceController) DeleteImage(ctx context.Context, eventId string) error {
	return s.repository.UpdateImage(ctx, eventId, "")
}
