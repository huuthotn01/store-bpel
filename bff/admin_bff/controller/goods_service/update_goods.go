package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) UpdateGoods(ctx context.Context, request []*goods_service.UpdateGoodsRequestData) error {
	goodsList := make([]*schema.UpdateGoodsRequest, 0, len(request))
	for _, goods := range request {
		goodsList = append(goodsList, &schema.UpdateGoodsRequest{
			GoodsSize:    goods.GoodsSize,
			GoodsColor:   goods.GoodsColor,
			GoodsName:    goods.GoodsName,
			GoodsType:    goods.GoodsType,
			GoodsGender:  goods.GoodsGender,
			GoodsAge:     goods.GoodsAge,
			Manufacturer: goods.Manufacturer,
			UnitPrice:    goods.UnitPrice,
			Description:  goods.Description,
		})
	}
	return c.goodsAdapter.UpdateGoods(ctx, request[0].GoodsCode, goodsList)

}
