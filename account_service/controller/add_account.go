package controller

import (
	"context"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) AddAccount(ctx context.Context, request *schema.AddAccountRequest) error {
	// only used in add staff account, staff service doesn't provide password
	// => account service generate random initial password for account
	rawPass := c.generateRandomPassword()
	hashedPass, err := c.hashPasswordBcrypt(rawPass)
	if err != nil {
		return err
	}

	return c.repository.AddAccount(ctx, &repository.AccountModel{
		Username: request.Username,
		Password: hashedPass,
		UserRole: request.Role,
	})
}
