package controller

import (
	"context"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error) {
	orders, err := c.repository.GetOnlineOrders(ctx, customerId)
	if err != nil {
		return nil, err
	}

	respOrders := make([]*schema.GetListOrderCustomerResponseData, 0, len(orders))
	for _, data := range orders {
		respOrders = append(respOrders, &schema.GetListOrderCustomerResponseData{
			OrderId:       data.OrderData.OrderCode,
			OrderCode:     data.OrderData.PublicOrderCode,
			PaymentMethod: data.OnlineOrderData.PaymentMethod,
			TotalPrice:    data.OrderData.TotalPrice,
		})
	}

	return respOrders, nil
}
