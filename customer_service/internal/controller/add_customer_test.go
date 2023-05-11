package controller

import (
	"context"
	"store-bpel/customer_service/schema"
	"testing"
)

func Test_customerServiceController_AddCustomer(t *testing.T) {
	type args struct {
		request *schema.AddCustomerRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add customer successfully",
			args: args{
				request: &schema.AddCustomerRequest{
					Username: "huutho",
					Email:    "huutho@gmail.com",
					Name:     "Huu Tho",
					Phone:    "0123456789",
					Gender:   "MALE",
					Age:      18,
					Street:   "LTK",
					Ward:     "P. 11",
					District: "Q. 10",
					Province: "TP. HCM",
				},
			},
		},
		{
			name: "Should return error when username existed",
			args: args{
				request: &schema.AddCustomerRequest{
					Username: "httn",
					Email:    "httn@gmail.com",
					Name:     "Huu Tho",
					Phone:    "0123456789",
					Gender:   "MALE",
					Age:      18,
					Street:   "LTK",
					Ward:     "P. 11",
					District: "Q. 10",
					Province: "TP. HCM",
				},
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
			if err := c.AddCustomer(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
