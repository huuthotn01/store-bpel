package cart_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/cart_service"
	"store-bpel/cart_service/schema"
)

func (c *cartBffController) UpdateGoods(ctx context.Context, request *cart_service.UpdateGoodsRequest) error {

	listGoods := make([]*schema.UpdateGoodsRequest, 0, len(request.Goods))

	for _, good := range request.Goods {
		listGoods = append(listGoods, &schema.UpdateGoodsRequest{
			GoodsId:    good.GoodsId,
			GoodsColor: good.GoodsColor,
			GoodsSize:  good.GoodsSize,
			Quantity:   good.Quantity,
		})
	}

	return c.cartAdapter.UpdateGoods(ctx, request.CartId, listGoods)
}
