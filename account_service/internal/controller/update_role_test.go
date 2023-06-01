package controller

import (
	"context"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_UpdateRole(t *testing.T) {
	type args struct {
		username string
		request  *schema.UpdateRoleRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update role successfully, role not 7",
			args: args{
				username: "test-user",
				request: &schema.UpdateRoleRequest{
					Role: 3,
				},
			},
		},
		{
			name: "Should return error when db return error getting account",
			args: args{
				username: "db-error-check-account",
				request: &schema.UpdateRoleRequest{
					Role: 3,
				},
			},
			wantErr: true,
		},
		{
			name: "Should update role successfully, role 7, request role not 7",
			args: args{
				username: "user-role-7",
				request: &schema.UpdateRoleRequest{
					Role: 3,
				},
			},
		},
		{
			name: "Should update role successfully, role 7, request role 7",
			args: args{
				username: "user-role-7",
				request: &schema.UpdateRoleRequest{
					Role: 7,
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
				repository: testRepository,
			}
			if err := c.UpdateRole(ctx, tt.args.username, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
