package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*schema.GetGoodsInWarehouseResponseData, error) {
	warehouses, err := c.repository.GetWarehouseByGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}
	result := make([]*schema.GetGoodsInWarehouseResponseData, 0, len(warehouses))
	for _, warehouse := range warehouses {
		result = append(result, &schema.GetGoodsInWarehouseResponseData{
			GoodsCode:   warehouse.GoodsCode,
			GoodsSize:   warehouse.GoodsSize,
			GoodsColor:  warehouse.GoodsColor,
			WhCode:      warehouse.WhCode,
			Quantity:    warehouse.Quantity,
			CreatedDate: warehouse.CreatedDate,
			UpdatedDate: warehouse.UpdatedDate,
		})
	}
	return result, nil
}
