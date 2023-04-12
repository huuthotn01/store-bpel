package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) GetOverallStat(ctx context.Context, request *schema.CommonGetStatisticRequest) ([]*schema.GetOverallStatisticResponseData, error) {
	stat, err := c.repository.GetOverallStat(ctx, request.Start, request.End, "", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	respStat := make([]*schema.GetOverallStatisticResponseData, 0, len(stat))
	for _, data := range stat {
		respStat = append(respStat, &schema.GetOverallStatisticResponseData{
			Revenue: data.Revenue,
			Profit:  data.Profit,
			Date:    data.Date,
		})
	}

	return respStat, nil
}
