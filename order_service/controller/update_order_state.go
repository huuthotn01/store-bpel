package controller

import (
	"context"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) UpdateOrderState(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error {
	return c.repository.UpdateOrderState(ctx, &repository.OrderStateModel{
		OrderCode: request.OrderId,
		State:     request.State,
	})
}
