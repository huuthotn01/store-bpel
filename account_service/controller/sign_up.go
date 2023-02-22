package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"store-bpel/account_service/repository"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) SignUp(ctx context.Context, request *schema.SignUpRequest) error {
	_, err := c.repository.GetAccount(ctx, request.Username)
	if err == nil {
		return errors.New("username existed")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPass, err := c.hashPasswordBcrypt(request.Password)
	if err != nil {
		return err
	}

	err = c.repository.AddAccount(ctx, &repository.AccountModel{
		Username: request.Username,
		Password: hashedPass,
		UserRole: request.Role,
	})
	if err != nil {
		return err
	}

	// TODO add data to customer service

	return nil
}
