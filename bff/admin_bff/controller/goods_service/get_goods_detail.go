package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

func (c *goodsBffController) GetGoodsDetail(ctx context.Context, request *goods_service.GetGoodsDetailRequest) (*goods_service.GetGoodsResponseData, error) {
	goods, err := c.goodsAdapter.GetGoodsDetail(ctx, request.GoodsCode)
	if err != nil {
		return nil, err
	}

	classify := make([]*goods_service.GetGoodsResponseData_Classify, 0, len(goods.Classify))
	for _, data := range goods.Classify {
		classify = append(classify, &goods_service.GetGoodsResponseData_Classify{
			Color: data.Color,
			Size:  data.Size,
		})
	}

	return &goods_service.GetGoodsResponseData{
		GoodsId:      goods.GoodsId,
		Classify:     classify,
		GoodsName:    goods.GoodsName,
		GoodsType:    goods.GoodsType,
		GoodsGender:  goods.GoodsGender,
		GoodsAge:     goods.GoodsAge,
		Manufacturer: goods.Manufacturer,
		IsForSale:    goods.IsForSale,
		UnitPrice:    goods.UnitPrice,
		UnitCost:     goods.UnitCost,
		Description:  goods.Description,
	}, nil
}
