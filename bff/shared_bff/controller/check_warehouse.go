package controller

import (
	"context"
	"store-bpel/bff/shared_bff/schema"
	goods_schema "store-bpel/goods_service/schema"
)

func (c *goodsBffController) CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error) {
	checkWHElements := make([]*goods_schema.CheckWarehouseRequestElement, 0, len(request.Elements))

	for _, data := range request.Elements {
		checkWHElements = append(checkWHElements, &goods_schema.CheckWarehouseRequestElement{
			GoodsCode:  data.GoodsCode,
			GoodsColor: data.GoodsColor,
			GoodsSize:  data.GoodsSize,
			Quantity:   data.Quantity,
		})
	}

	resp, err := c.goodsAdapter.CheckWarehouse(ctx, &goods_schema.CheckWarehouseRequest{
		Elements: checkWHElements,
	})
	if err != nil {
		return nil, err
	}

	whActions := make([]*schema.WarehouseActions, 0, len(resp.WarehouseActions))
	for _, action := range resp.WarehouseActions {
		whActions = append(whActions, &schema.WarehouseActions{
			GoodsCode:  action.GoodsCode,
			GoodsColor: action.GoodsColor,
			GoodsSize:  action.GoodsSize,
			Quantity:   action.Quantity,
			From:       action.From,
			To:         action.To,
		})
	}

	return &schema.CheckWarehouseResponseData{
		NeedTransfer:     resp.NeedTransfer,
		WarehouseActions: whActions,
	}, nil
}
