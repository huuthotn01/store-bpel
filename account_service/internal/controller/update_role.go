package controller

import (
	"context"
	"store-bpel/account_service/internal/util"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error {
	account, err := c.repository.GetAccount(ctx, username)
	if err != nil {
		return err
	}

	if account.UserRole == 7 {
		if request.Role != 7 {
			rawPass := util.GenerateRandomPassword()
			hashedPass, err := util.HashPasswordBcrypt(rawPass)
			if err != nil {
				return err
			}
			err = c.repository.UpdateRole(ctx, username, request.Role, hashedPass)
			to := []string{account.Email}
			title := "TÀI KHOẢN TRUY CẬP HỆ THỐNG PTH FASHION ADMIN"
			message := "Thông tin tài khoản của bạn:" +
				"\nUsername: " + account.Username +
				"\nPassword: " + rawPass

			return util.SendEmail(to, title, message)
		}
	}

	return c.repository.UpdateRole(ctx, username, request.Role, "")
}
