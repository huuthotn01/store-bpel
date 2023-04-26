package cart_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/cart_service"
	"store-bpel/cart_service/schema"
)

func (c *cartBffController) AddGoods(ctx context.Context, request *cart_service.AddGoodsRequest) error {

	listGoods := make([]*schema.AddGoodsRequest, 0, len(request.Goods))

	for _, good := range request.Goods {
		listGoods = append(listGoods, &schema.AddGoodsRequest{
			GoodsId:    good.GoodsId,
			GoodsColor: good.GoodsColor,
			GoodsSize:  good.GoodsSize,
			Quantity:   good.Quantity,
		})
	}

	return c.cartAdapter.AddGoods(ctx, request.CartId, listGoods)
}
