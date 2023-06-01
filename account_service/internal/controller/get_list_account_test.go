package controller

import (
	"context"
	"reflect"
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
		{
			name: "Should return error when get list account from db failed",
			args: args{
				username: "get-list-account-fail",
			},
			wantErr: true,
		},
		{
			name: "Should return error when customer adapter failed",
			args: args{
				username: "cust-adapter-fail",
			},
			wantErr: true,
		},
		{
			name: "Should return error when staff adapter failed",
			args: args{
				username: "staff-adapter-fail",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &accountServiceController{
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
