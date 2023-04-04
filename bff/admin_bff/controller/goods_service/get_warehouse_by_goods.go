package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

func (c *goodsBffController) GetWarehouseByGoods(ctx context.Context, request *goods_service.GetWarehouseByGoodsRequest) ([]*goods_service.GetWarehouseResponseData, error) {
	warehouses, err := c.goodsAdapter.GetWarehouseByGoods(ctx, request.GoodsCode)
	if err != nil {
		return nil, err
	}

	var result []*goods_service.GetWarehouseResponseData

	for _, warehouse := range warehouses {
		result = append(result, &goods_service.GetWarehouseResponseData{

			GoodsCode:  warehouse.GoodsCode,
			GoodsSize:  warehouse.GoodsSize,
			GoodsColor: warehouse.GoodsColor,
			WhCode:     warehouse.WhCode,
			Quantity:   warehouse.Quantity,
		})
	}

	return result, nil
}
