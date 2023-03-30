package controller

import (
	"context"
	"errors"
	"fmt"
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
	if len(goodsData) == 0 {
		return errors.New(fmt.Sprintf("no goods in any warehouse with goodsId %s, goodsColor %s, goodsSize %s and warehouse %s",
			request.GoodsCode, request.GoodsColor, request.GoodsSize, request.From))
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
