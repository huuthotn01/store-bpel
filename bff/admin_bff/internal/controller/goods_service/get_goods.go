package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

func (c *goodsBffController) GetGoods(ctx context.Context) ([]*goods_service.GetGoodsResponseData, error) {
	goods, err := c.goodsAdapter.GetGoods(ctx)
	if err != nil {
		return nil, err
	}
	respGoods := make([]*goods_service.GetGoodsResponseData, 0, len(goods))
	for _, data := range goods {
		classify := make([]*goods_service.GetGoodsResponseData_Classify, 0, len(data.Classify))
		for _, cl := range data.Classify {
			classify = append(classify, &goods_service.GetGoodsResponseData_Classify{
				Color: cl.Color,
				Size:  cl.Size,
			})
		}
		respGoods = append(respGoods, &goods_service.GetGoodsResponseData{
			GoodsId:      data.GoodsId,
			Classify:     classify,
			GoodsName:    data.GoodsName,
			GoodsType:    data.GoodsType,
			GoodsGender:  data.GoodsGender,
			GoodsAge:     data.GoodsAge,
			Manufacturer: data.Manufacturer,
			IsForSale:    data.IsForSale,
			UnitPrice:    data.UnitPrice,
			UnitCost:     data.UnitCost,
			Description:  data.Description,
			Image:        data.Image,
		})
	}

	return respGoods, nil
}
