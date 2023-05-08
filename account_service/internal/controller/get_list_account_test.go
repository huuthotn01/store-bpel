package controller

import (
	"context"
	"reflect"
	"store-bpel/account_service/config"
	"store-bpel/account_service/schema"
	"testing"
)

func Test_accountServiceController_GetListAccount(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetListAccountResponseData
		wantErr bool
	}{
		{
			name: "Should get list account",
			args: args{
				username: "",
			},
			want: []*schema.GetListAccountResponseData{
				{
					Username:    "cust-1",
					Id:          "cust-1",
					Role:        1,
					PhoneNumber: "0111111111",
					Email:       "cust-1@gmail.com",
					Name:        "Customer One",
					IsActivated: true,
				},
				{
					Username:    "staff-1@gmail.com",
					Id:          "staff-1",
					Role:        3,
					PhoneNumber: "0123456789",
					Email:       "staff-1@gmail.com",
					Name:        "Staff One",
					IsActivated: true,
				},
			},
		},
	}

	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
				cfg:             cfg,
				repository:      testRepository,
				staffAdapter:    testStaff,
				customerAdapter: testCustomer,
			}
			got, err := c.GetListAccount(ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
