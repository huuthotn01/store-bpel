package controller

import (
	"context"
	"reflect"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetOnlineOrdersStatus(t *testing.T) {
	type args struct {
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetOnlineOrdersStatusResponseData
		wantErr bool
	}{
		{
			name: "Should get online order status correctly",
			args: args{
				orderId: "order-1",
			},
			want: []*schema.GetOnlineOrdersStatusResponseData{
				{
					OrderId:   5,
					State:     "Packed by seller",
					StateTime: "2023-01-01 07:00:00",
				},
				{
					OrderId:   5,
					State:     "Picked from warehouse",
					StateTime: "2023-01-03 13:00:00",
				},
				{
					OrderId:   5,
					State:     "Delivered",
					StateTime: "2023-01-06 09:00:00",
				},
			},
		},
		{
			name: "Should return error when db get private code fails",
			args: args{
				orderId: "invalid-order",
			},
			wantErr: true,
		},
		{
			name: "Should return error when db get order state fails",
			args: args{
				orderId: "invalid-order-state",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{
				repository: testRepository,
			}
			got, err := c.GetOnlineOrdersStatus(ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOnlineOrdersStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOnlineOrdersStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
