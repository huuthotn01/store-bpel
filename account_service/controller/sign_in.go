package controller

import (
	"context"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) SignIn(ctx context.Context, request *schema.SignInRequest) (int, error) {
	account, err := c.repository.GetAccount(ctx, request.Username)
	if err != nil {
		return 0, err
	}
	err = c.checkPasswordBcrypt([]byte(account.Password), []byte(request.Password))
	if err != nil {
		return 0, err
	}
	return account.UserRole, nil
}
