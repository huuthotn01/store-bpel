package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) CheckWarehouse(ctx context.Context, request *goods_service.CheckWarehouseRequest) (*goods_service.CheckWarehouseResponseData, error) {
	checkWHElements := make([]*schema.CheckWarehouseRequestElement, 0, len(request.Elements))

	for _, data := range request.Elements {
		checkWHElements = append(checkWHElements, &schema.CheckWarehouseRequestElement{
			GoodsCode:  data.GoodsCode,
			GoodsColor: data.GoodsColor,
			GoodsSize:  data.GoodsSize,
			Quantity:   data.Quantity,
		})
	}

	resp, err := c.goodsAdapter.CheckWarehouse(ctx, &schema.CheckWarehouseRequest{
		Elements: checkWHElements,
	})
	if err != nil {
		return nil, err
	}

	whActions := make([]*goods_service.WarehouseActions, 0, len(resp.WarehouseActions))
	for _, action := range resp.WarehouseActions {
		whActions = append(whActions, &goods_service.WarehouseActions{
			GoodsCode:  action.GoodsCode,
			GoodsColor: action.GoodsColor,
			GoodsSize:  action.GoodsSize,
			Quantity:   action.Quantity,
			From:       action.From,
			To:         action.To,
		})
	}

	return &goods_service.CheckWarehouseResponseData{
		NeedTransfer:     resp.NeedTransfer,
		WarehouseActions: whActions,
	}, nil
}
