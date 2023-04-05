package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) AddGoods(ctx context.Context, request []*goods_service.AddGoodsRequestData) error {
	goodsList := make([]*schema.AddGoodsRequest, 0, len(request))
	for _, goods := range request {
		goodsList = append(goodsList, &schema.AddGoodsRequest{
			GoodsSize:    goods.GoodsSize,
			GoodsColor:   goods.GoodsColor,
			GoodsName:    goods.GoodsName,
			GoodsGender:  goods.GoodsGender,
			GoodsType:    goods.GoodsType,
			GoodsAge:     goods.GoodsAge,
			Manufacturer: goods.Manufacturer,
			IsForSale:    goods.IsForSale,
			UnitPrice:    goods.UnitPrice,
			Description:  goods.Description,
		})
	}

	return c.goodsAdapter.AddGoods(ctx, goodsList)
}
