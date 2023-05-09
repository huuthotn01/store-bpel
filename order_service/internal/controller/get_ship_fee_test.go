package controller

import (
	"context"
	"reflect"
	"store-bpel/order_service/schema"
	"testing"
	"time"
)

func Test_orderServiceController_GetShipFee(t *testing.T) {
	type args struct {
		request *schema.GetShipFeeRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetShipFeeResponseData
		wantErr bool
	}{
		{
			name: "Should get shipping fee and expected delivery correctly",
			args: args{
				request: &schema.GetShipFeeRequest{
					Street:   "LTK",
					Ward:     "Ward 11",
					District: "District 10",
					Province: "Ho Chi Minh City",
				},
			},
			want: &schema.GetShipFeeResponseData{
				ShipFee:      10000,
				ExpectedDate: time.Now().Local().Add(5 * 24 * time.Hour).Format("2006-01-02"),
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{}
			got, err := c.GetShipFee(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShipFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShipFee() got = %v, want %v", got, tt.want)
			}
		})
	}
}
