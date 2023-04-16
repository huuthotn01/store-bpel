package controller

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
	customer_schema "store-bpel/customer_service/schema"
	"store-bpel/library/kafka_lib"
)

func (c *accountServiceController) SignUp(ctx context.Context, request *schema.SignUpRequest) error {
	_, err := c.repository.GetAccount(ctx, request.Username)
	if err == nil {
		return errors.New("username existed")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPass, err := c.hashPasswordBcrypt(request.Password)
	if err != nil {
		return err
	}

	err = c.repository.AddAccount(ctx, &repository.AccountModel{
		Username: request.Username,
		Password: hashedPass,
		UserRole: 1,
	})
	if err != nil {
		return err
	}

	addCustomerRequest := &customer_schema.AddCustomerRequest{
		Username: request.Username,
		Email:    request.Email,
		Name:     request.Name,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Age:      request.Age,
		Street:   request.Street,
		Ward:     request.Ward,
		District: request.District,
		Province: request.Province,
	}

	if c.cfg.ServiceFlags.IsEnableAsync {
		addCustReqByte, err := json.Marshal(addCustomerRequest)
		if err != nil {
			return err
		}

		// publish event to customer service to add customer
		return c.kafkaAdapter.Publish(ctx, kafka_lib.CUSTOMER_SERVICE_TOPIC, addCustReqByte)
	}

	return c.customerAdapter.AddCustomer(ctx, addCustomerRequest)
}
