package controller

import (
	"context"
	"store-bpel/event_service/internal/repository"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) UpdateEvent(ctx context.Context, eventId string, request *schema.UpdateEventRequest) error {
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
