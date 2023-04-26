package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/order_service"
)

func (c *orderBffController) GetListOrderCustomer(ctx context.Context, request *order_service.GetListOrderCustomerRequest) ([]*order_service.GetListOrderCustomerResponseData, error) {
	orders, err := c.orderAdapter.GetListOrderCustomer(ctx, request.CustomerId)
	if err != nil {
		return nil, err
	}

	respOrders := make([]*order_service.GetListOrderCustomerResponseData, 0, len(orders))
	for _, order := range orders {
		listGoods := make([]*order_service.OrderGoodsResponse, 0, len(order.ListGoods))
		for _, goods := range order.ListGoods {
			listGoods = append(listGoods, &order_service.OrderGoodsResponse{
				GoodsId:   goods.GoodsId,
				Image:     goods.Image,
				Name:      goods.Name,
				UnitPrice: goods.UnitPrice,
				Price:     goods.Price,
				Quantity:  goods.Quantity,
				Size:      goods.Size,
				Color:     goods.Color,
				Discount:  goods.Discount,
			})
		}

		listState := make([]*order_service.GetListOrderStateResponse, 0, len(order.StatusShip))
		for _, state := range order.StatusShip {
			listState = append(listState, &order_service.GetListOrderStateResponse{
				State: state.State,
				Time:  state.Time,
			})
		}

		respOrders = append(respOrders, &order_service.GetListOrderCustomerResponseData{
			OrderId:         order.OrderId,
			OrderCode:       order.OrderCode,
			PaymentMethod:   order.PaymentMethod,
			ListGoods:       listGoods,
			TotalPrice:      order.TotalPrice,
			TotalGoods:      order.TotalGoods,
			TotalDiscount:   order.TotalDiscount,
			TotalOrder:      order.TotalOrder,
			IsCompleted:     order.IsCompleted,
			ShipFee:         order.ShipFee,
			StatusShip:      listState,
			TransactionDate: order.TransactionDate,
			ExpectDate:      order.ExpectDate,
		})
	}

	return respOrders, nil
}
