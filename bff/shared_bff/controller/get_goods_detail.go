package controller

import (
	"context"
	"store-bpel/bff/shared_bff/schema"
)

func (c *goodsBffController) GetGoodsDetail(ctx context.Context, request *schema.GetGoodsDetailRequest) (*schema.GetGoodsResponseData, error) {
	goods, err := c.goodsAdapter.GetGoodsDetail(ctx, request.GoodsCode)
	if err != nil {
		return nil, err
	}

	return &schema.GetGoodsResponseData{
		GoodsCode:    goods.GoodsCode,
		GoodsSize:    goods.GoodsSize,
		GoodsColor:   goods.GoodsColor,
		GoodsName:    goods.GoodsName,
		GoodsType:    goods.GoodsType,
		GoodsGender:  goods.GoodsGender,
		GoodsAge:     goods.GoodsAge,
		Manufacturer: goods.Manufacturer,
		IsForSale:    goods.IsForSale,
		UnitPrice:    goods.UnitPrice,
		Description:  goods.Description,
	}, nil
}
