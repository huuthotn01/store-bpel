package customer_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/customer_service"
)

func (c *customerBffController) GetCustomer(ctx context.Context, request *customer_service.GetCustomerInfoRequest) (*customer_service.GetCustomerInfoResponseData, error) {
	cust, err := c.customerAdapter.GetCustomer(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	return &customer_service.GetCustomerInfoResponseData{
		Username: cust.Username,
		Email:    cust.Email,
		Name:     cust.Name,
		Phone:    cust.Phone,
		Age:      cust.Age,
		Gender:   cust.Gender,
		Street:   cust.Street,
		Ward:     cust.Ward,
		District: cust.District,
		Province: cust.Province,
	}, nil
}
