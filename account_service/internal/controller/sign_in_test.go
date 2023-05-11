package controller

import (
	"context"
	"reflect"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_SignIn(t *testing.T) {
	type args struct {
		request *schema.SignInRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.SignInResponseData
		wantErr bool
	}{
		{
			name: "Should sign in successfully",
			args: args{
				request: &schema.SignInRequest{
					Username: "test-user",
					Password: "testpwd",
				},
			},
			want: &schema.SignInResponseData{
				UserId: "test-user",
				Role:   1,
			},
		},
		{
			name: "Should return error when account not activated",
			args: args{
				request: &schema.SignInRequest{
					Username: "unactivated",
					Password: "testpwd",
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
			got, err := c.SignIn(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				tt.want.Token = got.Token
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}
