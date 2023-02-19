package controller

import (
	"context"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) AddAccount(ctx context.Context, request *schema.AddAccountRequest) error {
	hashedPass, err := c.hashPasswordBcrypt(request.Password)
	if err != nil {
		return err
	}
	return c.repository.AddAccount(ctx, &repository.AccountModel{
		Username: request.Username,
		Password: hashedPass,
		UserRole: request.Role,
	})
}
