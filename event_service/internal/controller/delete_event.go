package controller

import (
	"context"
)

func (s *eventServiceController) DeleteEvent(ctx context.Context, eventId string) error {
	return s.repository.DeleteEvent(ctx, eventId)
}
