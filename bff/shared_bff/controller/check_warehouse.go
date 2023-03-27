package controller

import (
	"context"
	"store-bpel/bff/shared_bff/schema"
	goods_schema "store-bpel/goods_service/schema"
)

func (c *goodsBffController) CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error) {
	resp, err := c.goodsAdapter.CheckWarehouse(ctx, &goods_schema.CheckWarehouseRequest{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
		Quantity:   request.Quantity,
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
