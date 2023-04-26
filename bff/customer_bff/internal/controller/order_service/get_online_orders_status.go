package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/order_service"
)

func (c *orderBffController) GetOnlineOrdersStatus(ctx context.Context, request *order_service.GetOnlineOrdersStatusRequest) ([]*order_service.GetOnlineOrdersStatusResponseData, error) {
	status, err := c.orderAdapter.GetOnlineOrdersStatus(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	respStatus := make([]*order_service.GetOnlineOrdersStatusResponseData, 0, len(status))
	for _, data := range status {
		respStatus = append(respStatus, &order_service.GetOnlineOrdersStatusResponseData{
			OrderId:   data.OrderId,
			State:     data.State,
			StateTime: data.StateTime,
		})
	}

	return respStatus, nil
}
