package controller

import (
	"context"
	"fmt"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
	"time"
)

func (c *goodsServiceController) AddGoods(ctx context.Context, request []*schema.AddGoodsRequest) error {
	goodsId := fmt.Sprintf("goods_%d", time.Now().Unix())

	goodsList := make([]*repository.GoodsModel, 0, len(request))

	for _, goods := range request {
		goodsList = append(goodsList, &repository.GoodsModel{
			GoodsCode:    goodsId,
			GoodsSize:    goods.GoodsSize,
			GoodsColor:   goods.GoodsColor,
			GoodsName:    goods.GoodsName,
			GoodsType:    goods.GoodsType,
			GoodsGender:  goods.GoodsGender,
			GoodsAge:     goods.GoodsAge,
			Manufacturer: goods.Manufacturer,
			IsForSale:    1,
			UnitPrice:    goods.UnitPrice,
			UnitCost:     goods.UnitCost,
			Description:  goods.Description,
		})

	}

	return c.repository.AddGoods(ctx, goodsList)
}
