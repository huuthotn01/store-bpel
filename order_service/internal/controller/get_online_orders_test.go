package controller

import (
	"context"
	"reflect"
	"sort"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetOnlineOrders(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetOnlineOrdersResponseData
		wantErr bool
	}{
		{
			name: "Should get online orders correctly",
			want: []*schema.GetOnlineOrdersResponseData{
				{
					OrderId:         5,
					OrderCode:       "MNOPQR",
					TotalPrice:      40000,
					TotalOrder:      45000,
					TransactionDate: "2023-01-01",
					TotalDiscount:   8000,
					TotalGoods:      2,
					ListGoods: []*schema.OrderGoodsResponse{
						{
							GoodsId:   "goods-1",
							Name:      "Goods One",
							UnitPrice: 10000,
							Price:     20000,
							Quantity:  2,
							Size:      "XL",
							Color:     "red",
							Discount:  0.2,
						},
					},
					OnlineOrderData: &schema.OnlineOrderData{
						PaymentMethod: "COD",
						CustomerId:    "customer-1",
						ShipFee:       5000,
						ExpectDate:    "2023-01-06",
						Status:        3,
						NameReceiver:  "HTTN",
						PhoneReceiver: "0123456789",
						EmailReceiver: "httn@gmail.com",
						Address: &schema.Address{
							Street:   "THT",
							Ward:     "Ward 11",
							District: "District 10",
							Province: "Ho Chi Minh City",
						},
					},
				},
				{
					OrderId:         6,
					OrderCode:       "AAAAAA",
					TotalPrice:      22000,
					TotalOrder:      25000,
					TransactionDate: "2023-01-02",
					TotalDiscount:   4400,
					TotalGoods:      1,
					ListGoods: []*schema.OrderGoodsResponse{
						{
							GoodsId:   "goods-2",
							Name:      "Goods Two",
							UnitPrice: 20000,
							Price:     22000,
							Quantity:  1,
							Size:      "XXL",
							Color:     "blue",
							Discount:  0.2,
						},
					},
					OnlineOrderData: &schema.OnlineOrderData{
						PaymentMethod: "momo",
						CustomerId:    "customer-2",
						ShipFee:       3000,
						ExpectDate:    "2023-01-08",
						Status:        4,
						IsCompleted:   true,
						NameReceiver:  "Huu Tho",
						PhoneReceiver: "0111111111",
						EmailReceiver: "tho@gmail.com",
						Address: &schema.Address{
							Street:   "LTK",
							Ward:     "P.11",
							District: "Q.10",
							Province: "HCMC",
						},
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
			got, err := c.GetOnlineOrders(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOnlineOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].OrderId < got[j].OrderId
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOnlineOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}
