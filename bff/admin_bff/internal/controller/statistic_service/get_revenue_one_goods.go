package statistic_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/statistic_service"
	"store-bpel/statistic_service/schema"
)

func (c *statisticBffController) GetRevenueOneGoods(ctx context.Context, request *statistic_service.GetStatOneGoodsRequest) ([]*statistic_service.GetRevenueResponseData, error) {
	revenue, err := c.statisticAdapter.GetRevenueOneGoods(ctx, &schema.CommonGetStatisticRequest{
		Start: request.Start,
		End:   request.End,
	}, request.GoodsId)
	if err != nil {
		return nil, err
	}

	respRevenue := make([]*statistic_service.GetRevenueResponseData, 0, len(revenue))
	for _, data := range revenue {
		respRevenue = append(respRevenue, &statistic_service.GetRevenueResponseData{
			Revenue: data.Revenue,
			Date:    data.Date,
		})
	}

	return respRevenue, nil
}
