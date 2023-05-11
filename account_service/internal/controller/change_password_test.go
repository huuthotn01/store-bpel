package controller

import (
	"context"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_ChangePassword(t *testing.T) {
	type args struct {
		request *schema.ChangePasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should change password successfully",
			args: args{
				request: &schema.ChangePasswordRequest{
					Username:    "test-user",
					OldPassword: "testpwd",
					NewPassword: "newpwd",
				},
			},
		},
		{
			name: "Should return error old password not match",
			args: args{
				request: &schema.ChangePasswordRequest{
					Username:    "test-user",
					OldPassword: "unmatchedpwd",
					NewPassword: "newpwd",
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
			if err := c.ChangePassword(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
