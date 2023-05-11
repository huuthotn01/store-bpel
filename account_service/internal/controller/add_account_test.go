package controller

import (
	"context"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_AddAccount(t *testing.T) {
	type args struct {
		request *schema.AddAccountRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add account successfully, role not 7",
			args: args{
				request: &schema.AddAccountRequest{
					Username: "user-test",
					Role:     3,
					Email:    "test@mail.com",
				},
			},
		},
		{
			name: "Should add account successfully, role 7",
			args: args{
				request: &schema.AddAccountRequest{
					Username: "user-test",
					Role:     7,
					Email:    "test@mail.com",
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
			if err := c.AddAccount(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
