package controller

import (
	"context"
	"store-bpel/customer_service/schema"
)

func (c *customerServiceController) GetCustomerInfo(ctx context.Context, customerId string) (*schema.GetCustomerInfoResponseData, error) {
	customerData, err := c.repository.GetCustomerInfo(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return &schema.GetCustomerInfoResponseData{
		Username: customerData.Username,
		Email:    customerData.CustomerEmail,
		Name:     customerData.CustomerName,
		Phone:    customerData.CustomerPhone,
		Gender:   customerData.CustomerGender,
		Age:      customerData.CustomerAge,
		Street:   customerData.Street,
		Ward:     customerData.Ward,
		District: customerData.District,
		Province: customerData.Province,
	}, nil
}
