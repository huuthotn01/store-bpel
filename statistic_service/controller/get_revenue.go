package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) GetRevenue(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetRevenueResponseData, error) {
	stat, err := c.repository.GetOverallStat(ctx, request.Start, request.End, "", request.BranchId, request.Gender, request.Type)
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
