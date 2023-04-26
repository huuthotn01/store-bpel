package event_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/event_service"
)

func (c *eventBffController) GetEventCurrent(ctx context.Context, request *event_service.GetEventCurrentRequest) ([]*event_service.GetEventData, error) {
	// Call event adapter
	events, err := c.eventAdapter.GetEventCurrent(ctx, request.NextDate)
	if err != nil {
		return nil, err
	}

	var result []*event_service.GetEventData

	for _, event := range events {
		result = append(result, &event_service.GetEventData{
			Id:        event.Id,
			Name:      event.Name,
			Discount:  event.Discount,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			Image:     event.Image,
			Goods:     event.Goods,
		})
	}

	return result, nil

}
