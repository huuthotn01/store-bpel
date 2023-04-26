package controller

import (
	"context"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) GetEventDetail(ctx context.Context, eventId string) (*schema.GetEventData, error) {
	// call repository
	event, err := s.repository.GetEvent(ctx, eventId)

	// check error call repository
	if err != nil {
		return nil, err
	}

	goods, err := s.repository.GetGoods(ctx, event.EventId)
	if err != nil {
		return nil, err
	}

	result := &schema.GetEventData{
		Id:        event.EventId,
		Name:      event.Name,
		Discount:  event.Discount,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		Image:     event.Image,
		Goods:     goods,
	}

	return result, nil

}
