package controller

import (
	"context"
	"fmt"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
	"time"
)

func (c *goodsServiceController) AddGoods(ctx context.Context, request *schema.AddGoodsRequest) error {
	goodsId := fmt.Sprintf("goods_%d", time.Now().Unix())

	return c.repository.AddGoods(ctx, &repository.GoodsModel{
		GoodsCode:    goodsId,
		GoodsSize:    request.GoodsSize,
		GoodsColor:   request.GoodsColor,
		GoodsName:    request.GoodsName,
		GoodsType:    request.GoodsType,
		GoodsGender:  request.GoodsGender,
		GoodsAge:     request.GoodsAge,
		Manufacturer: request.Manufacturer,
		IsForSale:    1,
		UnitPrice:    request.UnitPrice,
		UnitCost:     request.UnitCost,
		Description:  request.Description,
	})
}
