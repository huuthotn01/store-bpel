package controller

import (
	"context"
	"fmt"
	repository2 "store-bpel/event_service/internal/repository"
	"store-bpel/event_service/schema"
	"time"
)

func (s *eventServiceController) AddEvent(ctx context.Context, request *schema.AddEventRequest) error {
	eventId := fmt.Sprintf("event_%d", time.Now().Unix())
	// call repository
	eventModel := &repository2.EventModel{
		EventId:   eventId,
		Name:      request.Name,
		Discount:  request.Discount,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Image:     request.Image,
	}
	return s.repository.AddEvent(ctx, &repository2.AddEventData{
		EventModel: eventModel,
		GoodsList:  request.Goods,
	})
}
