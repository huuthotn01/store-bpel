package controller

import (
	"context"
	"store-bpel/bff/shared_bff/schema"
	goods_schema "store-bpel/goods_service/schema"
)

func (c *goodsBffController) CreateTransfer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	return c.goodsAdapter.CreateWHTransfer(ctx, &goods_schema.CreateGoodsTransactionRequest{
		GoodsSize:  request.GoodsSize,
		GoodsColor: request.GoodsColor,
		GoodsCode:  request.GoodsCode,
		Quantity:   request.Quantity,
		From:       request.From,
		To:         request.To,
	})
}
