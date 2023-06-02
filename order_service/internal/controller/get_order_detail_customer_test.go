package controller

import (
	"context"
	"reflect"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetOrderDetail(t *testing.T) {
	type args struct {
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetOrderDetailCustomerResponseData
		wantErr bool
	}{
		{
			name: "Should get order detail successfully",
			args: args{
				orderId: "order-1",
			},
			want: &schema.GetOrderDetailCustomerResponseData{
				OrderCode: "ABCDEF",
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
				TotalPrice:      11000,
				TotalGoods:      4,
				TotalDiscount:   2000,
				TotalOrder:      12000,
				TransactionDate: "2023-01-01",
				PaymentMethod:   "COD",
				IsCompleted:     true,
				ShipFee:         1000,
				ExpectDate:      "2023-01-06",
				Status:          4,
				NameReceiver:    "HTTN",
				PhoneReceiver:   "0123456789",
				EmailReceiver:   "httn@gmail.com",
				Address: &schema.Address{
					Street:   "LTK",
					Ward:     "11",
					District: "10",
					Province: "HCMC",
				},
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
			},
		},
		{
			name: "Should return error when db get private code fails",
			args: args{
				orderId: "invalid-order",
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
			got, err := c.GetOrderDetail(ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
