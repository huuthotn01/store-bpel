package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) GetRevenueOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetRevenueResponseData, error) {
	stat, err := c.repository.GetOverallStat(ctx, request.Start, request.End, goodsId, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	respStat := make([]*schema.GetRevenueResponseData, 0, len(stat))
	for _, data := range respStat {
		respStat = append(respStat, &schema.GetRevenueResponseData{
			Revenue: data.Revenue,
			Date:    data.Date,
		})
	}

	return respStat, nil
}
