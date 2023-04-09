package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/order_service"
	"store-bpel/order_service/schema"
)

func (c *orderBffController) CreateOfflineOrder(ctx context.Context, request *order_service.MakeOfflineOrderRequest) error {
	coreGoods := make([]*schema.OrderGoodsRequest, 0, len(request.GoodsList))

	for _, data := range request.GoodsList {
		coreGoods = append(coreGoods, &schema.OrderGoodsRequest{
			GoodsId:   data.GoodsId,
			UnitPrice: data.UnitPrice,
			Price:     data.Price,
			Quantity:  data.Quantity,
			Size:      data.Size,
			Color:     data.Color,
			Discount:  data.Discount,
			Tax:       data.Tax,
		})
	}

	return c.orderAdapter.CreateOfflineOrders(ctx, &schema.MakeOfflineOrderRequest{
		TotalPrice:      request.TotalPrice,
		TransactionDate: request.TransactionDate,
		GoodsList:       coreGoods,
		StaffId:         request.StaffId,
		BranchId:        request.BranchId,
	})
}
