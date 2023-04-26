package cart_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/cart_service"
	"store-bpel/cart_service/schema"
)

func (c *cartBffController) DeleteGoods(ctx context.Context, request *cart_service.DeleteGoodsRequest) error {

	listGoods := make([]*schema.DeleteGoodsRequest, 0, len(request.Goods))

	for _, good := range request.Goods {
		listGoods = append(listGoods, &schema.DeleteGoodsRequest{
			GoodsId:    good.GoodsId,
			GoodsColor: good.GoodsColor,
			GoodsSize:  good.GoodsSize,
		})
	}

	return c.cartAdapter.DeleteGoods(ctx, request.CartId, listGoods)
}

func (c *cartBffController) DeleteAllGoods(ctx context.Context, cartId string) error {
	return c.cartAdapter.DeleteAllGoods(ctx, cartId)
}
