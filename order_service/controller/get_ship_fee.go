package controller

import (
	"context"
	"store-bpel/order_service/schema"
)

func (c *orderServiceController) GetShipFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error) {
	return &schema.GetShipFeeResponseData{
		ShipFee: 10000,
	}, nil
}
