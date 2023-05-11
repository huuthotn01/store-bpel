package controller

import (
	"context"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_SignUp(t *testing.T) {
	type args struct {
		request *schema.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should sign up successfully",
			args: args{
				request: &schema.SignUpRequest{
					Username: "new-customer",
					Password: "new-cust",
				},
			},
		},
		{
			name: "Should return error account existed",
			args: args{
				request: &schema.SignUpRequest{
					Username: "test-user",
					Password: "new-cust",
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
				repository:   testRepository,
				kafkaAdapter: testKafka,
			}
			if err := c.SignUp(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
