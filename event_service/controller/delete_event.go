package controller

import (
	"context"
)

func (s *eventServiceController) DeleteEvent(ctx context.Context, eventId int) error {
	err := s.repository.DeleteEvent(ctx, eventId)
	if err != nil {
		return err
	}

	return s.repository.DeleteGoods(ctx, eventId)

}
