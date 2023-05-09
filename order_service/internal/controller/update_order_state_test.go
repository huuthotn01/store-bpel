package controller

import (
	"context"
	"store-bpel/order_service/config"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_UpdateOrderState(t *testing.T) {
	type args struct {
		request *schema.UpdateOnlineOrdersStatusRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update order state successfully",
			args: args{
				request: &schema.UpdateOnlineOrdersStatusRequest{
					OrderId:      "order-1",
					State:        "DELIVERED",
					StatusNumber: 4,
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
			c := &orderServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := c.UpdateOrderState(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrderState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
