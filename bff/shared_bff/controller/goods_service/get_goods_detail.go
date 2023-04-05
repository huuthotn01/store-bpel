package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/goods_service"
)

func (c *goodsBffController) GetGoodsDetail(ctx context.Context, request *goods_service.GetGoodsDetailRequest) ([]*goods_service.GetGoodsResponseData, error) {
	goods, err := c.goodsAdapter.GetGoodsDetail(ctx, request.GoodsCode)
	if err != nil {
		return nil, err
	}

	respGoods := make([]*goods_service.GetGoodsResponseData, 0, len(goods))
	for _, data := range goods {
		respGoods = append(respGoods, &goods_service.GetGoodsResponseData{
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
			UnitCost:     data.UnitCost,
			Description:  data.Description,
		})
	}

	return respGoods, nil
}
