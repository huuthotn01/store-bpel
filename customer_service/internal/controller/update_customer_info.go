package controller

import (
	"context"
	"store-bpel/customer_service/internal/repository"
	"store-bpel/customer_service/schema"
)

func (c *customerServiceController) UpdateCustomerInfo(ctx context.Context, customerId string, request *schema.UpdateCustomerInfoRequest) error {
	return c.repository.UpdateCustomerInfo(ctx, &repository.CustomerModel{
		Username:       customerId,
		CustomerName:   request.Name,
		CustomerEmail:  request.Email,
		CustomerPhone:  request.Phone,
		CustomerAge:    request.Age,
		CustomerGender: request.Gender,
		Street:         request.Street,
		Ward:           request.Ward,
		District:       request.District,
		Province:       request.Province,
	})
}
