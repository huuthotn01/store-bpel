package controller

import (
	"context"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetOfflineOrders(ctx context.Context) ([]*schema.GetOfflineOrdersResponseData, error) {
	offlineOrders, err := c.repository.GetOfflineOrders(ctx)
	if err != nil {
		return nil, err
	}

	var (
		mapOrderIdToOrdersData = make(map[int]*schema.GetOfflineOrdersResponseData, 0)
		mapOrderIdToOrderModel = make(map[int][]*repository.GoodsModel, 0)
		respData               = make([]*schema.GetOfflineOrdersResponseData, 0, len(offlineOrders))
	)

	for _, order := range offlineOrders {
		// add order data if not exists
		if _, ok := mapOrderIdToOrdersData[order.OrderCode]; !ok {
			mapOrderIdToOrdersData[order.OrderCode] = &schema.GetOfflineOrdersResponseData{
				OrderId:         order.OrderCode,
				OrderCode:       order.PublicOrderCode,
				TotalPrice:      order.TotalPrice,
				TotalOrder:      order.TotalPrice, // offline order doesn't have shipping fee
				TransactionDate: order.TransactionDate,
				OfflineOrderData: &schema.OfflineOrderData{
					StaffId:  order.StaffId,
					BranchId: order.StoreCode,
				},
			}
		}

		// add order goods data
		orderGoodsList, ok := mapOrderIdToOrderModel[order.OrderCode]
		if !ok {
			// not exists goods list with this order id, create new slice
			mapOrderIdToOrderModel[order.OrderCode] = []*repository.GoodsModel{
				{
					GoodsCode:  order.GoodsCode,
					GoodsSize:  order.GoodsSize,
					GoodsColor: order.GoodsColor,
					GoodsName:  order.GoodsName,
					OrderCode:  order.OrderCode,
					Quantity:   order.Quantity,
					UnitPrice:  order.UnitPrice,
					TotalPrice: order.Price,
					Tax:        order.Tax,
					Image:      order.Image,
					Promotion:  order.Promotion,
				},
			}
		} else {
			// already exists, just append
			orderGoodsList = append(orderGoodsList, &repository.GoodsModel{
				GoodsCode:  order.GoodsCode,
				GoodsSize:  order.GoodsSize,
				GoodsColor: order.GoodsColor,
				GoodsName:  order.GoodsName,
				OrderCode:  order.OrderCode,
				Quantity:   order.Quantity,
				UnitPrice:  order.UnitPrice,
				TotalPrice: order.Price,
				Tax:        order.Tax,
				Image:      order.Image,
				Promotion:  order.Promotion,
			})
		}
	}

	// mapping goods model data
	for orderId, goods := range mapOrderIdToOrderModel {
		res := c.mapListGoods(goods)
		mapOrderIdToOrdersData[orderId].ListGoods = res.ListGoods
		mapOrderIdToOrdersData[orderId].TotalDiscount = res.TotalDiscount
		mapOrderIdToOrdersData[orderId].TotalGoods = res.TotalGoods
	}

	for _, data := range mapOrderIdToOrdersData {
		respData = append(respData, data)
	}

	return respData, nil
}
