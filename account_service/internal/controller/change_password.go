package controller

import (
	"context"
	"store-bpel/account_service/internal/util"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) ChangePassword(ctx context.Context, request *schema.ChangePasswordRequest) error {
	account, err := c.repository.GetAccount(ctx, request.Username)
	if err != nil {
		return err
	}

	err = util.CheckPasswordBcrypt([]byte(account.Password), []byte(request.OldPassword))
	if err != nil {
		return err
	}

	hashedPass, err := util.HashPasswordBcrypt(request.NewPassword)
	if err != nil {
		return err
	}

	return c.repository.UpdatePassword(ctx, request.Username, hashedPass)
}
