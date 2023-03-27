package controller

import (
	"context"
	"store-bpel/bff/shared_bff/schema"
)

func (c *goodsBffController) GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error) {
	goods, err := c.goodsAdapter.GetGoods(ctx)
	if err != nil {
		return nil, err
	}

	respGoods := make([]*schema.GetGoodsResponseData, 0, len(goods))
	for _, data := range goods {
		respGoods = append(respGoods, &schema.GetGoodsResponseData{
			GoodsCode:    data.GoodsCode,
			GoodsSize:    data.GoodsSize,
			GoodsColor:   data.GoodsColor,
			GoodsName:    data.GoodsName,
			GoodsType:    data.GoodsType,
			GoodsGender:  data.GoodsGender,
			GoodsAge:     data.GoodsAge,
			Manufacturer: data.Manufacturer,
			IsForSale:    data.IsForSale,
			UnitPrice:    data.UnitPrice,
			Description:  data.Description,
		})
	}

	return respGoods, nil
}
