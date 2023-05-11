package controller

import (
	"context"
	"reflect"
	"store-bpel/customer_service/schema"
	"testing"
)

func Test_customerServiceController_GetCustomerInfo(t *testing.T) {
	type args struct {
		customerId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetCustomerInfoResponseData
		wantErr bool
	}{
		{
			name: "Should get customer info correctly",
			args: args{
				customerId: "httn",
			},
			want: &schema.GetCustomerInfoResponseData{
				Username: "httn",
				Email:    "httn@gmail.com",
				Name:     "Huu Tho",
				Phone:    "0111111111",
				Gender:   "MALE",
				Street:   "THT",
				Ward:     "11",
				District: "10",
				Province: "HCMC",
			},
		},
		{
			name: "Should return error when customer id not existed",
			args: args{
				customerId: "huutho",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customerServiceController{
				repository: testRepository,
			}
			got, err := c.GetCustomerInfo(ctx, tt.args.customerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
