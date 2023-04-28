package controller

import (
	"context"
	"store-bpel/account_service/internal/repository"
	"store-bpel/account_service/internal/util"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) AddAccount(ctx context.Context, request *schema.AddAccountRequest) error {
	// only used in add staff account, staff service doesn't provide password
	// => account service generate random initial password for account
	rawPass := util.GenerateRandomPassword()
	hashedPass, err := util.HashPasswordBcrypt(rawPass)
	if err != nil {
		return err
	}

	err = c.repository.AddAccount(ctx, &repository.AccountModel{
		Username: request.Username,
		Password: hashedPass,
		Email:    request.Email,
		UserRole: request.Role,
	})

	if err != nil {
		return err
	}

	if request.Role != 7 {
		to := []string{request.Email}
		title := "TÀI KHOẢN TRUY CẬP HỆ THỐNG PTH FASHION ADMIN"
		message := "Thông tin tài khoản của bạn:" +
			"\nUsername: " + request.Username +
			"\nPassword: " + rawPass

		return util.SendEmail(to, title, message)
	}
	return nil
}
