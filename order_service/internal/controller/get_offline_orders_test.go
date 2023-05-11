package controller

import (
	"context"
	"reflect"
	"sort"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_GetOfflineOrders(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetOfflineOrdersResponseData
		wantErr bool
	}{
		{
			name: "Should get offline orders correctly",
			want: []*schema.GetOfflineOrdersResponseData{
				{
					OrderId:         3,
					OrderCode:       "MNOPQR",
					TotalPrice:      40000,
					TotalOrder:      40000, // offline order doesn't have shipping fee
					TransactionDate: "2023-01-01",
					TotalDiscount:   8000,
					TotalGoods:      2,
					ListGoods: []*schema.OrderGoodsResponse{
						{
							GoodsId:   "goods-11",
							Name:      "Goods Eleven",
							UnitPrice: 10000,
							Price:     20000,
							Quantity:  2,
							Size:      "XL",
							Color:     "red",
							Discount:  0.2,
						},
					},
					OfflineOrderData: &schema.OfflineOrderData{
						StaffId:  "staff-1",
						BranchId: "store-1",
					},
				},
				{
					OrderId:         4,
					OrderCode:       "AAAAAA",
					TotalPrice:      22000,
					TotalOrder:      22000, // offline order doesn't have shipping fee
					TransactionDate: "2023-01-02",
					TotalGoods:      1,
					TotalDiscount:   4400,
					ListGoods: []*schema.OrderGoodsResponse{
						{
							GoodsId:   "goods-12",
							Name:      "Goods Twelve",
							UnitPrice: 20000,
							Price:     22000,
							Quantity:  1,
							Size:      "XXL",
							Color:     "blue",
							Discount:  0.2,
						},
					},
					OfflineOrderData: &schema.OfflineOrderData{
						StaffId:  "staff-3",
						BranchId: "store-2",
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
			got, err := c.GetOfflineOrders(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOfflineOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].OrderId < got[j].OrderId
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOfflineOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}
