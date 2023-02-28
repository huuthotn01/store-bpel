package controller

import (
	"context"
	"errors"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) handleWHTransfer(ctx context.Context, request *schema.CreateGoodsTransactionRequest) error {
	goodsData, err := c.repository.GetGoodsInWHData(ctx, &repository.GoodsInWh{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
		WhCode:     request.From,
	})
	if err != nil {
		return err
	}
	if goodsData[0].Quantity-request.Quantity < 0 {
		return errors.New("given quantity is greater than current quantity")
	}

	return c.repository.UpdateGoodsInWHTransfer(ctx, &repository.GoodsInWh{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
		Quantity:   request.Quantity,
	}, request.From, request.To)
}
