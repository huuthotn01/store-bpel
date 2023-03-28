package controller

import (
	"context"
	"store-bpel/event_service/repository"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) UpdateEvent(ctx context.Context, eventId int, request *schema.UpdateEventRequest) error {
	// call repository
	eventModel := &repository.EventModel{
		Name:      request.Name,
		Discount:  request.Discount,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Image:     request.Image,
	}

	return s.repository.UpdateEvent(ctx, &repository.UpdateEventData{
		EventModel: eventModel,
		EventId:    eventId,
		GoodsList:  request.Goods,
	})

}
