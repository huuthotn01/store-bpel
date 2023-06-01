package controller

import (
	"context"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_CreateResetPassword(t *testing.T) {
	type args struct {
		request *schema.CreateResetPasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create reset password successfully",
			args: args{
				request: &schema.CreateResetPasswordRequest{
					Username: "test-user",
				},
			},
		},
		{
			name: "Should return error when db return error getting account",
			args: args{
				request: &schema.CreateResetPasswordRequest{
					Username: "db-error-check-account",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when db return error updating otp code",
			args: args{
				request: &schema.CreateResetPasswordRequest{
					Username: "db-update-otp-fail",
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
				repository: testRepository,
			}
			if err := c.CreateResetPassword(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("CreateResetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_accountServiceController_ConfirmOTP(t *testing.T) {
	type args struct {
		request *schema.ConfirmOTPRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should confirm OTP and gen new password successfully",
			args: args{
				request: &schema.ConfirmOTPRequest{
					Username: "test-user",
					Otp:      "test-otp",
				},
			},
		},
		{
			name: "Should return error when OTP not correct",
			args: args{
				request: &schema.ConfirmOTPRequest{
					Username: "test-user",
					Otp:      "invalid-otp",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when db return error updating password",
			args: args{
				request: &schema.ConfirmOTPRequest{
					Username: "db-update-password-fail",
					Otp:      "test-otp",
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
				repository: testRepository,
			}
			if err := c.ConfirmOTP(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("ConfirmOTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
