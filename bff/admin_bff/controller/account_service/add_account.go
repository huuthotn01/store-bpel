package account_service

import (
	"context"
	"store-bpel/account_service/schema"
	"store-bpel/bff/admin_bff/schema/account_service"
)

func (c *accountBffController) AddAccount(ctx context.Context, request *account_service.AddAccountRequest) error {
	return c.accountAdapter.AddAccount(ctx, &schema.AddAccountRequest{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	})
}
