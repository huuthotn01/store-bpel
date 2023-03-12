package controller

import (
	"context"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error) {
	account, err := c.repository.GetAccount(ctx, request.Username)
	if err != nil {
		return nil, err
	}
	err = c.checkPasswordBcrypt([]byte(account.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}
	return &schema.SignInResponseData{
		Role: account.UserRole,
	}, nil
}
