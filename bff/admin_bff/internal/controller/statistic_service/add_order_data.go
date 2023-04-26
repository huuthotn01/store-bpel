package statistic_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/statistic_service"
	"store-bpel/statistic_service/schema"
)

func (c *statisticBffController) AddOrderData(ctx context.Context, request *statistic_service.AddOrderDataRequest) error {
	goodsData := make([]*schema.AddOrderDataRequest_GoodsData, 0, len(request.GoodsData))
	for _, data := range goodsData {
		goodsData = append(goodsData, &schema.AddOrderDataRequest_GoodsData{
			GoodsId:     data.GoodsId,
			GoodsSize:   data.GoodsSize,
			GoodsColor:  data.GoodsColor,
			GoodsType:   data.GoodsType,
			GoodsGender: data.GoodsGender,
			GoodsCost:   data.GoodsCost,
			UnitPrice:   data.UnitPrice,
			Quantity:    data.Quantity,
		})
	}

	return c.statisticAdapter.AddOrderData(ctx, &schema.AddOrderDataRequest{
		OrderId:         request.OrderId,
		TransactionDate: request.TransactionDate,
		ShopCode:        request.ShopCode,
		GoodsData:       goodsData,
	})
}
