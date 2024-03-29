package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/statistic_service/config"
	repo "store-bpel/statistic_service/internal/repository"
	"store-bpel/statistic_service/schema"
)

type IStatisticServiceController interface {
	GetOverallStat(ctx context.Context, request *schema.CommonGetStatisticRequest) ([]*schema.GetOverallStatisticResponseData, error)
	GetRevenue(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetRevenueResponseData, error)
	GetRevenueOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetRevenueResponseData, error)
	GetProfit(ctx context.Context, request *schema.FilterGetStatisticRequest) ([]*schema.GetProfitResponseData, error)
	GetProfitOneGoods(ctx context.Context, request *schema.CommonGetStatisticRequest, goodsId string) ([]*schema.GetProfitResponseData, error)
	AddOrderData(ctx context.Context, request *schema.AddOrderDataRequest) error
}

type statisticServiceController struct {
	cfg        *config.Config
	repository repo.IStatisticServiceRepository
}

func NewController(cfg *config.Config, db *gorm.DB) IStatisticServiceController {
	// init repository
	repository := repo.NewRepository(db)

	return &statisticServiceController{
		cfg:        cfg,
		repository: repository,
	}
}
