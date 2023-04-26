package controller

import (
	"context"
	"store-bpel/goods_service/internal/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) UpdateGoods(ctx context.Context, request []*schema.UpdateGoodsRequest, goodsId string) error {
	goodsList := make([]*repository.GoodsModel, 0, len(request))
	for _, item := range request {
		goodsList = append(goodsList, &repository.GoodsModel{
			GoodsCode:    goodsId,
			GoodsAge:     item.GoodsAge,
			GoodsGender:  item.GoodsGender,
			GoodsType:    item.GoodsType,
			GoodsColor:   item.GoodsColor,
			GoodsSize:    item.GoodsSize,
			GoodsName:    item.GoodsName,
			Manufacturer: item.Manufacturer,
			UnitPrice:    item.UnitPrice,
			IsForSale:    1,
			UnitCost:     item.UnitCost,
			Description:  item.Description,
		})
	}

	return c.repository.UpdateGoods(ctx, goodsList)
}
