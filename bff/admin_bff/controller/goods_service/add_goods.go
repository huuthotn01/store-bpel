package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) AddGoods(ctx context.Context, request *goods_service.AddGoodsRequest) error {
	return c.goodsAdapter.AddGoods(ctx, &schema.AddGoodsRequest{
		GoodsSize:    request.GoodsSize,
		GoodsColor:   request.GoodsColor,
		GoodsName:    request.GoodsName,
		GoodsGender:  request.GoodsGender,
		GoodsType:    request.GoodsType,
		GoodsAge:     request.GoodsAge,
		Manufacturer: request.Manufacturer,
		IsForSale:    request.IsForSale,
		UnitPrice:    request.UnitPrice,
		Description:  request.Description,
	})
}
