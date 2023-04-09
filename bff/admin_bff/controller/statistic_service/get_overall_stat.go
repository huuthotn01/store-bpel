package statistic_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/statistic_service"
	"store-bpel/statistic_service/schema"
)

func (c *statisticBffController) GetOverallStat(ctx context.Context, request *statistic_service.GetOverallStatRequest) ([]*statistic_service.GetOverallStatisticResponseData, error) {
	stat, err := c.statisticAdapter.GetOverallStat(ctx, &schema.CommonGetStatisticRequest{
		Start: request.Start,
		End:   request.End,
	})
	if err != nil {
		return nil, err
	}

	respStat := make([]*statistic_service.GetOverallStatisticResponseData, 0, len(stat))
	for _, data := range stat {
		respStat = append(respStat, &statistic_service.GetOverallStatisticResponseData{
			Revenue: data.Revenue,
			Profit:  data.Profit,
			Date:    data.Date,
		})
	}

	return respStat, nil
}
