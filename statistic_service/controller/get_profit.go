package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) GetProfit(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetProfitResponseData, error) {
	stat, err := c.repository.GetOverallStat(ctx, request.Start, request.End, "", request.BranchId, request.Gender, request.Type)
	if err != nil {
		return nil, err
	}

	respStat := make([]*schema.GetProfitResponseData, 0, len(stat))
	for _, data := range respStat {
		respStat = append(respStat, &schema.GetProfitResponseData{
			Profit: data.Profit,
			Date:   data.Date,
		})
	}

	return respStat, nil
}
