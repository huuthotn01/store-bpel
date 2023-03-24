package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/order_service"
	"store-bpel/order_service/schema"
)

func (c *orderBffController) UpdateOnlineOrdersStatus(ctx context.Context, request *order_service.UpdateOnlineOrdersStatusRequest) error {
	return c.orderAdapter.UpdateOnlineOrdersStatus(ctx, &schema.UpdateOnlineOrdersStatusRequest{
		OrderId: request.OrderId,
		State:   request.State,
	})
}
