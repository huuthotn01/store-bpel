package controller

import (
	"context"
)

func (s *eventServiceController) DeleteEvent(ctx context.Context, eventId int) error {
	return s.repository.DeleteEvent(ctx, eventId)
}
