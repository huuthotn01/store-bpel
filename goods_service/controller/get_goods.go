package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error) {
	var (
		res              = make([]*schema.GetGoodsResponseData, 0)
		mapGoodsIdToData = make(map[string]*schema.GetGoodsResponseData, 0)
		mapGoodsIdToImg  = make(map[string][]string, 0)
	)

	images, err := c.repository.GetImages(ctx)
	if err != nil {
		return nil, err
	}
	for _, data := range images {
		if _, ok := mapGoodsIdToImg[data.GoodsCode]; ok {
			mapGoodsIdToImg[data.GoodsCode] = append(mapGoodsIdToImg[data.GoodsCode], data.GoodsImg)
		} else {
			mapGoodsIdToImg[data.GoodsCode] = []string{data.GoodsImg}
		}
	}

	goods, err := c.repository.GetGoods(ctx)
	if err != nil {
		return nil, err
	}

	for _, data := range goods {
		// already have data with goodsId, just append another classify
		if _, ok := mapGoodsIdToData[data.GoodsCode]; ok {
			mapGoodsIdToData[data.GoodsCode].Classify = append(mapGoodsIdToData[data.GoodsCode].Classify, &schema.GetGoodsResponseData_Classify{
				Color: data.GoodsColor,
				Size:  data.GoodsSize,
			})
			continue
		}

		// goodsId has not existed, init a new one in map
		tmp := &schema.GetGoodsResponseData{
			GoodsId:   data.GoodsCode,
			GoodsName: data.GoodsName,
			Classify: []*schema.GetGoodsResponseData_Classify{
				{
					Color: data.GoodsColor,
					Size:  data.GoodsSize,
				},
			},
			GoodsType:    data.GoodsType,
			GoodsGender:  data.GoodsGender,
			GoodsAge:     data.GoodsAge,
			Manufacturer: data.Manufacturer,
			IsForSale:    data.IsForSale,
			UnitCost:     data.UnitCost,
			UnitPrice:    data.UnitPrice,
			Description:  data.Description,
		}
		mapGoodsIdToData[data.GoodsCode] = tmp
	}

	for goodsId, data := range mapGoodsIdToData {
		data.Image = mapGoodsIdToImg[goodsId]
		res = append(res, data)
	}

	//// TODO handle WH call
	//_, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	//if err != nil {
	//	return nil, err
	//}
	return res, nil
}
