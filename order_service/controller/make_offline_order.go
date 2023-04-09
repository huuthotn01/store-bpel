package controller

import (
	"context"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) CreateOfflineOrder(ctx context.Context, request *schema.MakeOfflineOrderRequest) error {
	orderPublicCode := c.generateOrderPublicCode()

	orderGoods := make([]*repository.GoodsModel, 0, len(request.GoodsList))
	for _, data := range request.GoodsList {
		orderGoods = append(orderGoods, &repository.GoodsModel{
			GoodsCode:  data.GoodsId,
			GoodsSize:  data.Size,
			GoodsColor: data.Color,
			Quantity:   data.Quantity,
			UnitPrice:  data.UnitPrice,
			TotalPrice: data.Price,
			Tax:        data.Tax,
			GoodsName:  "",
			Image:      "", // TODO add image link
			Promotion:  data.Discount,
		})
	}

	return c.repository.CreateOfflineOrder(ctx, &repository.OfflineOrdersData{
		PublicOrderCode: orderPublicCode,
		TransactionDate: request.TransactionDate,
		TotalPrice:      request.TotalPrice,
		BranchId:        request.BranchId,
		StaffId:         request.StaffId,
		Goods:           orderGoods,
	})
}
