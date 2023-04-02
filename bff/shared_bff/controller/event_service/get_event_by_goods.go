package event_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/event_service"
)

func (c *eventBffController) GetEventByGoods(ctx context.Context, request *event_service.GetEventByGoodsRequest) ([]*event_service.GetEventByGoodsData, error) {
	// Call event adapter
	events, err := c.eventAdapter.GetEventByGoods(ctx, request.GoodsId)
	if err != nil {
		return nil, err
	}

	var result []*event_service.GetEventByGoodsData

	for _, event := range events {
		result = append(result, &event_service.GetEventByGoodsData{
			Id:        event.Id,
			Name:      event.Name,
			Discount:  event.Discount,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			Image:     event.Image,
		})
	}

	return result, nil

}
