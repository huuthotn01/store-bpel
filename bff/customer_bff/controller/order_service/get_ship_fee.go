package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/order_service"
	"store-bpel/order_service/schema"
)

func (c *orderBffController) GetShippingFee(ctx context.Context, request *order_service.Address) (int, error) {
	shipFee, err := c.orderAdapter.GetShippingFee(ctx, &schema.GetShipFeeRequest{
		Street:   request.Street,
		Ward:     request.Ward,
		District: request.District,
		Province: request.Province,
	})
	if err != nil {
		return 0, err
	}

	return shipFee.ShipFee, nil
}
