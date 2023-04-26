package customer_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/customer_service"
	"store-bpel/customer_service/schema"
)

func (c *customerBffController) UpdateCustomer(ctx context.Context, request *customer_service.UpdateCustomerInfoRequest) error {
	return c.customerAdapter.UpdateCustomer(ctx, request.Username, &schema.UpdateCustomerInfoRequest{
		Email:    request.Email,
		Name:     request.Name,
		Age:      request.Age,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Street:   request.Street,
		Ward:     request.Ward,
		District: request.District,
		Province: request.Province,
	})
}
