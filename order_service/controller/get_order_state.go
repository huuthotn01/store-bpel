package controller

import (
	"context"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetOnlineOrdersStatus(ctx context.Context, orderId string) ([]*schema.GetOnlineOrdersStatusResponseData, error) {
	privateOrderId, err := c.repository.GetPrivateOrderCode(ctx, orderId)
	if err != nil {
		return nil, err
	}

	status, err := c.repository.GetOrderState(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	respStatus := make([]*schema.GetOnlineOrdersStatusResponseData, 0, len(status))
	for _, data := range status {
		respStatus = append(respStatus, &schema.GetOnlineOrdersStatusResponseData{
			OrderId:   data.OrderCode,
			State:     data.State,
			StateTime: data.StateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return respStatus, nil
}
