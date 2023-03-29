package controller

import (
	"context"
	"store-bpel/order_service/schema"
	"time"
)

func (c *orderServiceController) GetShipFee(ctx context.Context, request *schema.GetShipFeeRequest) (*schema.GetShipFeeResponseData, error) {
	today := time.Now().Local().Format("2006-01-02")
	expectedDay, err := c.getExpectedDelivery(today)
	if err != nil {
		return nil, err
	}

	return &schema.GetShipFeeResponseData{
		ShipFee:      10000,
		ExpectedDate: expectedDay,
	}, nil
}

func (c *orderServiceController) getExpectedDelivery(transactionDate string) (string, error) {
	createdDate, err := time.Parse("2006-01-02", transactionDate)
	if err != nil {
		return "", err
	}

	// after 5 days
	return createdDate.Add(5 * 24 * time.Hour).Format("2006-01-02"), nil
}
