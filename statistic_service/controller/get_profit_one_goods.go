package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) GetProfitOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetProfitResponseData, error) {
	stat, err := c.repository.GetOverallStat(ctx, request.Start, request.End, goodsId, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	respStat := make([]*schema.GetProfitResponseData, 0, len(stat))
	for _, data := range stat {
		respStat = append(respStat, &schema.GetProfitResponseData{
			Profit: data.Profit,
			Date:   data.Date,
		})
	}

	return respStat, nil
}
