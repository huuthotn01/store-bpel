package customer_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/customer_service"
	"store-bpel/customer_service/schema"
)

func (c *customerBffController) AddCustomer(ctx context.Context, request *customer_service.AddCustomerRequest) error {
	return c.customerAdapter.AddCustomer(ctx, &schema.AddCustomerRequest{
		Email:    request.Email,
		Username: request.Username,
		Name:     request.Name,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Age:      request.Age,
		Street:   request.Street,
		Ward:     request.Ward,
		District: request.District,
		Province: request.Province,
	})
}
