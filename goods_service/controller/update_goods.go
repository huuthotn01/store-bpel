package controller

import (
	"context"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) UpdateGoods(ctx context.Context, request *schema.UpdateGoodsRequest, goodsId string) error {
	return c.repository.UpdateGoods(ctx, &repository.GoodsModel{
		GoodsCode:    goodsId,
		GoodsAge:     request.GoodsAge,
		GoodsGender:  request.GoodsGender,
		GoodsType:    request.GoodsType,
		GoodsColor:   request.GoodsColor,
		GoodsSize:    request.GoodsSize,
		GoodsName:    request.GoodsName,
		Manufacturer: request.Manufacturer,
		UnitPrice:    request.UnitPrice,
		Description:  request.Description,
	})
}
