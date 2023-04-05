package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetGoods(ctx context.Context) ([]*schema.GetGoodsResponseData, error) {
	goods, err := c.repository.GetGoods(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetGoodsResponseData, 0, len(goods))
	for _, item := range goods {
		converted := schema.GetGoodsResponseData(*item)
		res = append(res, &converted)
	}
	// TODO handle WH call
	_, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *goodsServiceController) GetDetailGoods(ctx context.Context, goodsId string) ([]*schema.GetGoodsResponseData, error) {
	goods, err := c.repository.GetDetailGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	res := make([]*schema.GetGoodsResponseData, 0, len(goods))
	for _, item := range goods {
		converted := schema.GetGoodsResponseData(*item)
		res = append(res, &converted)
	}
	// TODO handle WH call
	_, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	if err != nil {
		return nil, err
	}
	// resConverted := schema.GetGoodsResponseData(*goods)
	// // TODO handle WH call
	// _, err = c.warehouseServiceAdapter.GetWarehouse(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	return res, nil
}
