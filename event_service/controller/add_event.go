package controller

import (
	"context"
	"store-bpel/event_service/repository"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) AddEvent(ctx context.Context, data *schema.AddEventRequest) error {
	// call repository
	eventModel := &repository.EventModel{
		Name:      data.Name,
		Discount:  data.Discount,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Image:     data.Image,
	}

	eventId, err := s.repository.AddEvent(ctx, eventModel)

	// check error call repository
	if err != nil {
		return err
	}

	err = s.repository.AddGoods(ctx, eventId, data.Goods)
	return err
}
