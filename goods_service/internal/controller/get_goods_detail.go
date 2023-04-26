package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetDetailGoods(ctx context.Context, goodsId string) (*schema.GetGoodsResponseData, error) {
	goods, err := c.repository.GetDetailGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	images, err := c.repository.GetGoodsImageUrls(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	classify := make([]*schema.GetGoodsResponseData_Classify, 0, len(goods))
	for _, data := range goods {
		classify = append(classify, &schema.GetGoodsResponseData_Classify{
			Size:  data.GoodsSize,
			Color: data.GoodsColor,
		})
	}

	res := &schema.GetGoodsResponseData{
		GoodsId:      goodsId,
		GoodsName:    goods[0].GoodsName,
		Classify:     classify,
		GoodsGender:  goods[0].GoodsGender,
		GoodsAge:     goods[0].GoodsAge,
		GoodsType:    goods[0].GoodsType,
		Manufacturer: goods[0].Manufacturer,
		IsForSale:    goods[0].IsForSale,
		UnitPrice:    goods[0].UnitPrice,
		UnitCost:     goods[0].UnitCost,
		Description:  goods[0].Description,
		Image:        images,
	}

	//// TODO handle WH call
	//_, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//resConverted := schema.GetGoodsResponseData(*goods)
	//// TODO handle WH call
	//_, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	//if err != nil {
	//	return nil, err
	//}
	return res, nil
}
