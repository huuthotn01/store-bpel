package account_service

import (
	"context"
	"store-bpel/account_service/schema"
	"store-bpel/bff/admin_bff/schema/account_service"
)

func (c *accountBffController) SignIn(ctx context.Context, request *account_service.SignInRequest) (*account_service.SignInResponseData, error) {
	signInResp, err := c.accountAdapter.SignIn(ctx, &schema.SignInRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &account_service.SignInResponseData{
		UserId: signInResp.UserId,
		Role:   signInResp.Role,
	}, nil
}
