package controller

import (
	"context"
	"store-bpel/event_service/schema"
)

func (s *eventServiceController) GetEventByGoods(ctx context.Context, goodsId string) ([]*schema.GetEventByGoodsData, error) {
	events, err := s.repository.GetEventByGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	var result []*schema.GetEventByGoodsData
	for _, event := range events {
		result = append(result, &schema.GetEventByGoodsData{
			Id:        event.EventId,
			Name:      event.Name,
			Discount:  event.Discount,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			Image:     event.Image,
		})
	}

	return result, nil

}
