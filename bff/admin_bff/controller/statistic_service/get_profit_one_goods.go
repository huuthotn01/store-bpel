package statistic_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/statistic_service"
	"store-bpel/statistic_service/schema"
)

func (c *statisticBffController) GetProfitOneGoods(ctx context.Context, request *statistic_service.GetStatOneGoodsRequest) ([]*statistic_service.GetProfitResponseData, error) {
	profit, err := c.statisticAdapter.GetProfitOneGoods(ctx, &schema.CommonGetStatisticRequest{
		Start: request.Start,
		End:   request.End,
	}, request.GoodsId)
	if err != nil {
		return nil, err
	}

	respProfit := make([]*statistic_service.GetProfitResponseData, 0, len(profit))
	for _, data := range profit {
		respProfit = append(respProfit, &statistic_service.GetProfitResponseData{
			Profit: data.Profit,
			Date:   data.Date,
		})
	}

	return respProfit, nil
}
