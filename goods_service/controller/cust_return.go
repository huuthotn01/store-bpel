package controller

import (
	"context"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) handleCustReturn(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	return c.repository.UpdateGoodsInWHInOut(ctx, &repository.GoodsInWh{
		GoodsCode:  request.GoodsCode,
		GoodsSize:  request.GoodsSize,
		GoodsColor: request.GoodsColor,
		WhCode:     request.To,
		Quantity:   request.Quantity,
	})
}
