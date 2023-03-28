package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) CustomerReturn(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error {
	return c.goodsAdapter.CustomerReturn(ctx, &schema.CreateGoodsTransactionRequest{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
		Quantity:   request.Quantity,
		From:       request.From,
		To:         request.To,
	})
}
