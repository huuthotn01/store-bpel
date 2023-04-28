package account_service

import (
	"context"
	"store-bpel/account_service/schema"
	"store-bpel/bff/customer_bff/schema/account_service"
)

func (c *accountBffController) ChangePassword(ctx context.Context, request *account_service.ChangePasswordRequest) error {
	return c.accountAdapter.ChangePassword(ctx, &schema.ChangePasswordRequest{
		Username:    request.Username,
		OldPassword: request.OldPassword,
		NewPassword: request.NewPassword,
	})
}
