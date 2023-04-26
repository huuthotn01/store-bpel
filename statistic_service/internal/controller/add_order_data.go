package controller

import (
	"context"
	"store-bpel/statistic_service/internal/repository"
	"store-bpel/statistic_service/schema"
)

func (c *statisticServiceController) AddOrderData(ctx context.Context, request *schema.AddOrderDataRequest) error {
	orderModel := &repository.OrdersModel{
		OrderCode:       request.OrderId,
		TransactionDate: request.TransactionDate,
		ShopCode:        request.ShopCode,
	}

	goodsModel := make([]*repository.GoodsModel, 0, len(request.GoodsData))
	for _, data := range request.GoodsData {
		goodsModel = append(goodsModel, &repository.GoodsModel{
			GoodsCode:   data.GoodsId,
			GoodsSize:   data.GoodsSize,
			GoodsColor:  data.GoodsColor,
			GoodsType:   data.GoodsType,
			GoodsGender: data.GoodsGender,
			GoodsCost:   data.GoodsCost,
			UnitPrice:   data.UnitPrice,
			Quantity:    data.Quantity,
			OrderCode:   request.OrderId,
		})
	}

	return c.repository.AddOrderData(ctx, goodsModel, orderModel)
}
