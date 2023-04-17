package controller

import (
	"context"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) GetEventCurrent(ctx context.Context, date int) ([]*schema.GetEventData, error) {
	// call repository
	events, err := s.repository.GetAllEventCurrent(ctx, date)

	// check error call repository
	if err != nil {
		return nil, err
	}

	// khai báo slice
	res := make([]*schema.GetEventData, 0, len(events))

	for _, event := range events {

		goods, err := s.repository.GetGoods(ctx, event.EventId)
		if err != nil {
			return nil, err
		}

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
