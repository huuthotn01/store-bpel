package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) CreateTransfer(ctx context.Context, request *goods_service.CreateGoodsTransactionRequest) error {
	return c.goodsAdapter.CreateWHTransfer(ctx, &schema.CreateGoodsTransactionRequest{
		GoodsSize:  request.GoodsSize,
		GoodsColor: request.GoodsColor,
		GoodsCode:  request.GoodsCode,
		Quantity:   request.Quantity,
		From:       request.From,
		To:         request.To,
	})
}
