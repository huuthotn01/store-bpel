package controller

import (
	"context"
	"store-bpel/customer_service/config"
	"store-bpel/customer_service/schema"
	"testing"
)

func Test_customerServiceController_UpdateCustomerInfo(t *testing.T) {
	type args struct {
		customerId string
		request    *schema.UpdateCustomerInfoRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update customer info correctly",
			args: args{
				customerId: "httn",
				request: &schema.UpdateCustomerInfoRequest{
					Email: "httn1@gmail.com",
					Name:  "HTTN",
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
			c := &customerServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := c.UpdateCustomerInfo(ctx, tt.args.customerId, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
