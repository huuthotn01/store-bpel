package controller

import (
	"context"
	"store-bpel/event_service/repository"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) UpdateEvent(ctx context.Context, eventId int, data *schema.UpdateEventRequest) error {
	// call repository
	eventModel := &repository.EventModel{
		Name:      data.Name,
		Discount:  data.Discount,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Image:     data.Image,
	}

	err := s.repository.UpdateEvent(ctx, eventId, eventModel)

	// check error call repository
	if err != nil {
		return err
	}

	err = s.repository.DeleteGoods(ctx, eventId)
	if err != nil {
		return err
	}

	return s.repository.AddGoods(ctx, eventId, data.Goods)

}
