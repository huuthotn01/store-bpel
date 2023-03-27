package controller

import (
	"context"
	"errors"
	"sort"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error) {
	requestQuantity := request.Quantity

	whData, err := c.repository.GetGoodsInWHData(ctx, &repository.GoodsInWh{
		GoodsCode:  request.GoodsCode,
		GoodsColor: request.GoodsColor,
		GoodsSize:  request.GoodsSize,
	})
	if err != nil {
		return nil, err
	}

	mapIntWhData, quantities := c.convertWHDataToMap(whData)
	if quantities[0] >= requestQuantity {
		return &schema.CheckWarehouseResponseData{
			NeedTransfer: false,
		}, nil
	}
	destWarehouse := mapIntWhData[quantities[0]].WhCode
	requestQuantity -= quantities[0]
	quantities = quantities[1:]

	whActions := make([]*schema.WarehouseActions, 0)
	for requestQuantity > 0 {
		if len(quantities) == 0 {
			return nil, errors.New("requested amount exceeds total amount in all warehouses")
		}
		if quantities[0] >= requestQuantity {
			requestQuantity = 0
			whActions = append(whActions, &schema.WarehouseActions{
				GoodsCode:  request.GoodsCode,
				GoodsColor: request.GoodsColor,
				GoodsSize:  request.GoodsSize,
				From:       mapIntWhData[quantities[0]].WhCode,
				To:         destWarehouse,
				Quantity:   requestQuantity,
			})
		} else {
			requestQuantity -= quantities[0]
			whActions = append(whActions, &schema.WarehouseActions{
				GoodsCode:  request.GoodsCode,
				GoodsColor: request.GoodsColor,
				GoodsSize:  request.GoodsSize,
				From:       mapIntWhData[quantities[0]].WhCode,
				To:         destWarehouse,
				Quantity:   quantities[0],
			})
			quantities = quantities[1:]
		}
	}

	return &schema.CheckWarehouseResponseData{
		NeedTransfer:     true,
		WarehouseActions: whActions,
	}, nil
}

func (c *goodsServiceController) convertWHDataToMap(data []*repository.GoodsInWh) (map[int]*repository.GoodsInWh, []int) {
	resp := make(map[int]*repository.GoodsInWh, 0)
	quantities := make([]int, 0, len(data))
	for _, d := range data {
		quantities = append(quantities, d.Quantity)
		resp[d.Quantity] = d
	}

	// sort quantity in descending order
	sort.Slice(quantities, func(i, j int) bool {
		return quantities[i] > quantities[j]
	})

	return resp, quantities
}
