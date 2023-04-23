package controller

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"store-bpel/goods_service/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) CheckWarehouse(ctx context.Context, request *schema.CheckWarehouseRequest) (*schema.CheckWarehouseResponseData, error) {
	var (
		whActions     = make([]*schema.WarehouseActions, 0)
		checkWHResult = &schema.CheckWarehouseResponseData{
			NeedTransfer: false, // initially, no need for transferring
		}
	)

	for _, ele := range request.Elements {
		action, err := c.processCheckWarehouse(ctx, ele)
		if err != nil {
			return nil, err
		}

		if len(action) > 0 {
			checkWHResult.NeedTransfer = true
			whActions = append(whActions, action...)
		}
	}

	checkWHResult.WarehouseActions = whActions

	return checkWHResult, nil
}

func (c *goodsServiceController) processCheckWarehouse(ctx context.Context, data *schema.CheckWarehouseRequestElement) ([]*schema.WarehouseActions, error) {
	requestQuantity := data.Quantity

	whData, err := c.repository.GetGoodsInWHData(ctx, &repository.GoodsInWh{
		GoodsCode:  data.GoodsCode,
		GoodsColor: data.GoodsColor,
		GoodsSize:  data.GoodsSize,
	})
	if err != nil {
		return nil, err
	}
	if len(whData) == 0 {
		return nil, errors.New(fmt.Sprintf("not found goods in warehouse with GoodsId %s, GoodsColor %s, GoodsSize %s",
			data.GoodsCode, data.GoodsColor, data.GoodsSize))
	}

	mapIntWhData, quantities := c.convertWHDataToMap(whData)
	if quantities[0] >= requestQuantity {
		return []*schema.WarehouseActions{}, nil
	}
	destWarehouse := mapIntWhData[quantities[0]][0].WhCode
	requestQuantity -= quantities[0]
	mapIntWhData[quantities[0]] = mapIntWhData[quantities[0]][1:]
	quantities = quantities[1:]

	whActions := make([]*schema.WarehouseActions, 0)
	for requestQuantity > 0 {
		if len(quantities) == 0 {
			return nil, errors.New("requested amount exceeds total amount in all warehouses")
		}
		if quantities[0] >= requestQuantity {
			whActions = append(whActions, &schema.WarehouseActions{
				GoodsCode:  data.GoodsCode,
				GoodsColor: data.GoodsColor,
				GoodsSize:  data.GoodsSize,
				From:       mapIntWhData[quantities[0]][0].WhCode,
				To:         destWarehouse,
				Quantity:   requestQuantity,
			})
			requestQuantity = 0
		} else {
			requestQuantity -= quantities[0]
			whActions = append(whActions, &schema.WarehouseActions{
				GoodsCode:  data.GoodsCode,
				GoodsColor: data.GoodsColor,
				GoodsSize:  data.GoodsSize,
				From:       mapIntWhData[quantities[0]][0].WhCode,
				To:         destWarehouse,
				Quantity:   quantities[0],
			})
			mapIntWhData[quantities[0]] = mapIntWhData[quantities[0]][1:]
			quantities = quantities[1:]
		}
	}

	return whActions, nil
}

func (c *goodsServiceController) convertWHDataToMap(data []*repository.GoodsInWh) (map[int][]*repository.GoodsInWh, []int) {
	resp := make(map[int][]*repository.GoodsInWh, 0)
	quantities := make([]int, 0, len(data))
	for _, d := range data {
		quantities = append(quantities, d.Quantity)
		resp[d.Quantity] = append(resp[d.Quantity], d)
	}

	// sort quantity in descending order
	sort.Slice(quantities, func(i, j int) bool {
		return quantities[i] > quantities[j]
	})

	return resp, quantities
}
