package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/order_service"
)

func (c *orderBffController) GetOrderCustomerDetail(ctx context.Context, request *order_service.GetOrderDetailCustomerRequest) (*order_service.GetOrderDetailCustomerResponseData, error) {
	order, err := c.orderAdapter.GetOrderCustomerDetail(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

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

	return &order_service.GetOrderDetailCustomerResponseData{
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
		Status:          order.Status,
		NameReceiver:    order.NameReceiver,
		PhoneReceiver:   order.PhoneReceiver,
		EmailReceiver:   order.EmailReceiver,
		Address: &order_service.Address{
			Street:   order.Address.Street,
			Ward:     order.Address.Ward,
			District: order.Address.District,
			Province: order.Address.Province,
		},
		ExpectDate: order.ExpectDate,
	}, nil
}
