package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/order_service"
)

func (c *orderBffController) GetOnlineOrders(ctx context.Context) ([]*order_service.GetOnlineOrdersResponseData, error) {
	orders, err := c.orderAdapter.GetOnlineOrders(ctx)
	if err != nil {
		return nil, err
	}

	respOrders := make([]*order_service.GetOnlineOrdersResponseData, 0, len(orders))
	for _, order := range orders {
		listGoods        := make([]*order_service.OrderGoodsResponse, 0, len(order.ListGoods))
		for _, goods := range order.ListGoods {
			listGoods = append(listGoods, &order_service.OrderGoodsResponse{
				GoodsId:   goods.GoodsId,
				Image:     goods.Image,
				Name:      goods.Name,
				UnitPrice: goods.UnitPrice,
				Price:     goods.Price,
				Tax:       goods.Tax,
				Quantity:  goods.Quantity,
				Size:      goods.Size,
				Color:     goods.Color,
				Discount:  goods.Discount,
			})
		}

		respOrders = append(respOrders, &order_service.GetOnlineOrdersResponseData{
			OrderId:          order.OrderId,
			OrderCode:        order.OrderCode,
			ListGoods:        listGoods,
			TotalPrice:       order.TotalPrice,
			TotalGoods:       order.TotalGoods,
			TotalDiscount:    order.TotalDiscount,
			TotalOrder:       order.TotalOrder,
			TransactionDate:  order.TransactionDate,
			OnlineOrderData:  &order_service.OnlineOrderData{
				PaymentMethod: order.OnlineOrderData.PaymentMethod,
				CustomerId:    order.OnlineOrderData.CustomerId,
				IsCompleted:   order.OnlineOrderData.IsCompleted,
				ShipFee:       order.OnlineOrderData.ShipFee,
				ExpectDate:    order.OnlineOrderData.ExpectDate,
				Status:        order.OnlineOrderData.Status,
				NameReceiver:  order.OnlineOrderData.NameReceiver,
				PhoneReceiver: order.OnlineOrderData.PhoneReceiver,
				EmailReceiver: order.OnlineOrderData.EmailReceiver,
				Address: &order_service.Address{
					Street:   order.OnlineOrderData.Address.Street,
					Ward:     order.OnlineOrderData.Address.Ward,
					District: order.OnlineOrderData.Address.District,
					Province: order.OnlineOrderData.Address.Province,
				},
			},
		})
	}

	return respOrders, nil
}
