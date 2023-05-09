package controller

import (
	"context"
	"store-bpel/order_service/internal/repository"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) UpdateOrderState(ctx context.Context, request *schema.UpdateOnlineOrdersStatusRequest) error {
	privateOrderId, err := c.repository.GetPrivateOrderCode(ctx, request.OrderId)
	if err != nil {
		return err
	}

	return c.repository.UpdateOrderState(ctx, &repository.OnlineOrderStateData{
		OrderState: &repository.OrderStateModel{
			OrderCode: privateOrderId,
			State:     request.State,
		},
		StatusNumber: request.StatusNumber,
	})
}
