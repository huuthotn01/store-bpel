package account_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/account_service"
)

func (c *accountBffController) SignIn(ctx context.Context, request *account_service.SignInRequest) (*account_service.SignInResponseData, error) {
	return &account_service.SignInResponseData{
		Role: 2,
	}, nil
}
