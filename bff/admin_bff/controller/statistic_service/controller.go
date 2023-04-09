package statistic_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/statistic_service"
)

type IStatisticBffController interface {
	GetOverallStat(ctx context.Context, request *statistic_service.GetOverallStatRequest) ([]*statistic_service.GetOverallStatisticResponseData, error)
	GetRevenue(ctx context.Context, request *statistic_service.FilterGetStatisticRequest) ([]*statistic_service.GetRevenueResponseData, error)
	GetRevenueOneGoods(ctx context.Context, request *statistic_service.GetStatOneGoodsRequest) ([]*statistic_service.GetRevenueResponseData, error)
	GetProfit(ctx context.Context, request *statistic_service.FilterGetStatisticRequest) ([]*statistic_service.GetProfitResponseData, error)
	GetProfitOneGoods(ctx context.Context, request *statistic_service.GetStatOneGoodsRequest) ([]*statistic_service.GetProfitResponseData, error)
}

type statisticBffController struct {
	cfg              *config.Config
	statisticAdapter adapter.IStatisticServiceAdapter
}

func NewController(cfg *config.Config) IStatisticBffController {
	// init stat adapter
	statAdapter := adapter.NewStatisticAdapter(cfg)

	return &statisticBffController{
		cfg:              cfg,
		statisticAdapter: statAdapter,
	}
}
