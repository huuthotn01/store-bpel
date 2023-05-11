package controller

import (
	"context"
	"reflect"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetOrderDetailAdmin(t *testing.T) {
	type args struct {
		orderId int
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetOrderDetailAdminResponseData
		wantErr bool
	}{
		{
			name: "Should get order detail admin correctly, offline order",
			args: args{
				orderId: 3,
			},
			want: &schema.GetOrderDetailAdminResponseData{
				OrderId:   3,
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
				TotalDiscount:   2200,
				TotalOrder:      11000,
				TransactionDate: "2023-01-01",
				IsOnline:        false,
				OfflineOrderData: &schema.OfflineOrderData{
					StaffId:  "staff-1",
					BranchId: "store-1",
				},
			},
		},
		{
			name: "Should get order detail admin correctly, online order",
			args: args{
				orderId: 5,
			},
			want: &schema.GetOrderDetailAdminResponseData{
				OrderId:   5,
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
				TotalDiscount:   2200,
				TotalOrder:      12000,
				TransactionDate: "2023-01-01",
				IsOnline:        true,
				OnlineOrderData: &schema.OnlineOrderData{
					PaymentMethod: "COD",
					CustomerId:    "customer-1",
					IsCompleted:   true,
					ShipFee:       1000,
					ExpectDate:    "2023-01-06",
					Status:        4,
					NameReceiver:  "HTTN",
					PhoneReceiver: "0123456789",
					EmailReceiver: "httn@gmail.com",
					Address: &schema.Address{
						Street:   "LTK",
						Ward:     "11",
						District: "10",
						Province: "HCMC",
					},
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
			got, err := c.GetOrderDetailAdmin(ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderDetailAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrderDetailAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
