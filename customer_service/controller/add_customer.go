package controller

import (
	"context"
	"errors"
	"fmt"
	"store-bpel/customer_service/repository"
	"store-bpel/customer_service/schema"

	"gorm.io/gorm"
)

func (c *customerServiceController) AddCustomer(ctx context.Context, request *schema.AddCustomerRequest) error {
	_, err := c.GetCustomerInfo(ctx, request.Username)
	if err == nil {
		return errors.New(fmt.Sprintf("customer id %s existed", request.Username))
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	err = c.cartAdapter.AddCart(ctx, request.Username)
	if err != nil {
		return err
	}

	return c.repository.AddCustomer(ctx, &repository.CustomerModel{
		Username:       request.Username,
		CustomerEmail:  request.Email,
		CustomerName:   request.Name,
		CustomerPhone:  request.Phone,
		CustomerAge:    request.Age,
		CustomerGender: request.Gender,
		Street:         request.Street,
		Ward:           request.Ward,
		District:       request.District,
		Province:       request.Province,
	})
}
