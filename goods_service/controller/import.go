package controller

import (
	"context"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) handleImport(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	return c.repository.UpdateGoodsInWHInOut(ctx, &repository.GoodsInWh{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
		WhCode:     request.To,
		Quantity:   request.Quantity,
	})
}
