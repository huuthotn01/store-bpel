package controller

import (
	"context"
	"reflect"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetListOrderCustomer(t *testing.T) {
	type args struct {
		customerId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetListOrderCustomerResponseData
		wantErr bool
	}{
		{
			name: "Should get list order by customer correctly",
			args: args{
				customerId: "customer-1",
			},
			want: []*schema.GetListOrderCustomerResponseData{
				{
					OrderCode:     "ABCDEF",
					PaymentMethod: "COD",
					ListGoods: []*schema.OrderGoodsResponse{
						{
							GoodsId:   "goods-1",
							Name:      "Goods One",
							UnitPrice: 5000,
							Price:     5500,
							Quantity:  3,
							Size:      "XL",
							Color:     "red",
							Discount:  0.1,
						},
						{
							GoodsId:   "goods-2",
							Name:      "Goods Two",
							UnitPrice: 5000,
							Price:     5500,
							Quantity:  1,
							Size:      "S",
							Color:     "yellow",
							Discount:  0.1,
						},
					},
					TotalPrice:    11000,
					TotalGoods:    4,
					TotalDiscount: 2000,
					TotalOrder:    12000,
					IsCompleted:   true,
					ShipFee:       1000,
					StatusShip: []*schema.GetListOrderStateResponse{
						{
							State: "Packed by seller",
							Time:  "2023-01-01 15:16:17",
						},
						{
							State: "Picked by shipper",
							Time:  "2023-01-02 08:09:10",
						},
					},
					TransactionDate: "2023-01-01",
					ExpectDate:      "2023-01-06",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{
				repository: testRepository,
			}
			got, err := c.GetListOrderCustomer(ctx, tt.args.customerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListOrderCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListOrderCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
