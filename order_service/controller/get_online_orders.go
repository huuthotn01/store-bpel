package controller

import (
	"context"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetOnlineOrders(ctx context.Context) ([]*schema.GetOnlineOrdersResponseData, error) {
	onlineOrders, err := c.repository.GetOnlineOrders(ctx)
	if err != nil {
		return nil, err
	}

	var (
		mapOrderIdToOrdersData = make(map[int]*schema.GetOnlineOrdersResponseData, 0)
		mapOrderIdToOrderModel = make(map[int][]*repository.GoodsModel, 0)
		respData               = make([]*schema.GetOnlineOrdersResponseData, 0, len(onlineOrders))
	)

	for _, order := range onlineOrders {
		// add order data if not exists
		if _, ok := mapOrderIdToOrdersData[order.OrderCode]; !ok {
			mapOrderIdToOrdersData[order.OrderCode] = &schema.GetOnlineOrdersResponseData{
				OrderId:         order.OrderCode,
				OrderCode:       order.PublicOrderCode,
				TotalPrice:      order.TotalPrice,
				TotalOrder:      order.TotalPrice, // add shipping fee later
				TransactionDate: order.TransactionDate,
				OnlineOrderData: &schema.OnlineOrderData{
					PaymentMethod: order.PaymentMethod,
					CustomerId:    order.CustomerId,
					IsCompleted:   order.Status == 4,
					ShipFee:       order.ShippingFee,
					ExpectDate:    order.ExpectedDelivery,
					Status:        order.Status,
					NameReceiver:  order.CustomerName,
					PhoneReceiver: order.CustomerPhone,
					EmailReceiver: order.CustomerEmail,
					Address: &schema.Address{
						Street:   order.Street,
						Ward:     order.Ward,
						District: order.District,
						Province: order.Province,
					},
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
		mapOrderIdToOrdersData[orderId].TotalOrder = mapOrderIdToOrdersData[orderId].TotalOrder + mapOrderIdToOrdersData[orderId].OnlineOrderData.ShipFee
	}

	for _, data := range mapOrderIdToOrdersData {
		respData = append(respData, data)
	}

	return respData, nil
}
