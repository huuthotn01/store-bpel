package controller

import (
	"context"
	"store-bpel/event_service/repository"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) AddEvent(ctx context.Context, request *schema.AddEventRequest) error {
	// call repository
	eventModel := &repository.EventModel{
		Name:      request.Name,
		Discount:  request.Discount,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Image:     request.Image,
	}
	return s.repository.AddEvent(ctx, &repository.AddEventData{
		EventModel: eventModel,
		GoodsList:  request.Goods,
	})
}
