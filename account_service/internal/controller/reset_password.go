package controller

import (
	"context"
	"errors"
	"store-bpel/account_service/internal/util"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) CreateResetPassword(ctx context.Context, request *schema.CreateResetPasswordRequest) error {
	account, err := c.repository.GetAccount(ctx, request.Username)
	if err != nil {
		return err
	}

	otpCode := util.GenerateOTPCode()
	err = c.repository.UpdateOTPCode(ctx, request.Username, otpCode)
	if err != nil {
		return err
	}

	to := []string{account.Email}
	title := "MÃ XÁC THỰC PTH FASHION"
	message := "Lưu ý: Không tiết lộ mã OTP với bất kì ai, mã OTP có thời hạn 5 phút:" +
		"\nMã OTP: " + otpCode

	util.SendEmail(to, title, message)
	return nil
}

func (c *accountServiceController) ConfirmOTP(ctx context.Context, request *schema.ConfirmOTPRequest) error {
	account, err := c.repository.ConfirmOTP(ctx, request.Username, request.Otp)
	if err != nil {
		return errors.New("OTP is not correct")
	}

	rawPass := util.GenerateRandomPassword()
	hashedPass, err := util.HashPasswordBcrypt(rawPass)
	if err != nil {
		return err
	}

	err = c.repository.UpdatePassword(ctx, request.Username, hashedPass)

	if err != nil {
		return err
	}

	to := []string{account.Email}
	title := "MẬT KHẨU MỚI TRUY CẬP HỆ THỐNG PTH FASHION ADMIN"
	message := "Thông tin tài khoản của bạn:" +
		"\nUsername: " + request.Username +
		"\nPassword: " + rawPass

	return util.SendEmail(to, title, message)
}
