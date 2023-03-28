package controller

import (
	"context"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) GetEvent(ctx context.Context) ([]*schema.GetEventData, error) {
	// call repository
	events, err := s.repository.GetAllEvent(ctx)

	// check error call repository
	if err != nil {
		return nil, err
	}

	// khai báo slice
	res := make([]*schema.GetEventData, 0, len(events))
	goods := []*string{}

	for _, event := range events {
		res = append(res, &schema.GetEventData{
			Id:        event.EventId,
			Name:      event.Name,
			Discount:  event.Discount,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			Image:     event.Image,
			Goods:     goods,
		})
	}

	return res, nil

}
