package account_service

import (
	"context"
	"store-bpel/account_service/schema"
	"store-bpel/bff/shared_bff/schema/account_service"
)

func (c *accountBffController) CreateResetPassword(ctx context.Context, request *account_service.CreateResetPasswordRequest) error {
	return c.accountAdapter.CreateResetPassword(ctx, &schema.CreateResetPasswordRequest{
		Username: request.Username,
	})
}

func (c *accountBffController) ConfirmOTP(ctx context.Context, request *account_service.ConfirmOTPRequest) error {
	return c.accountAdapter.ConfirmOTP(ctx, &schema.ConfirmOTPRequest{
		Username: request.Username,
		Otp:      request.Otp,
	})
}
