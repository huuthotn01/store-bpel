package controller

import (
	"context"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetListOrderCustomer(ctx context.Context, customerId string) ([]*schema.GetListOrderCustomerResponseData, error) {
	orders, err := c.repository.GetOnlineOrdersByCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}

	respOrders := make([]*schema.GetListOrderCustomerResponseData, 0, len(orders))
	for _, data := range orders {
		mappingData := c.mapOrderDetailData(data)

		respOrders = append(respOrders, &schema.GetListOrderCustomerResponseData{
			// OrderId is private information => not return for customer
			OrderCode:       data.OrderData.PublicOrderCode,
			PaymentMethod:   data.OnlineOrderData.PaymentMethod,
			ListGoods:       mappingData.ListGoods,
			TotalPrice:      data.OrderData.TotalPrice,
			TotalGoods:      mappingData.GoodsNum,
			TotalDiscount:   mappingData.TotalDiscount,
			TotalOrder:      mappingData.TotalOrder,
			IsCompleted:     data.OnlineOrderData.Status == 4,
			ShipFee:         data.OnlineOrderData.ShippingFee,
			StatusShip:      mappingData.StatusShip,
			TransactionDate: data.OrderData.TransactionDate,
			ExpectDate:      data.OnlineOrderData.ExpectedDelivery,
		})
	}

	return respOrders, nil
}
