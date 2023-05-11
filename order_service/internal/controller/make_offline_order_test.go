package controller

import (
	"context"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_CreateOfflineOrder(t *testing.T) {
	type args struct {
		request *schema.MakeOfflineOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create offline order successfully",
			args: args{
				request: &schema.MakeOfflineOrderRequest{
					GoodsList: []*schema.OrderGoodsRequest{
						{
							GoodsId:   "goods-1",
							Size:      "XL",
							Color:     "red",
							Quantity:  3,
							UnitPrice: 5000,
							Price:     15000,
							Tax:       0.1,
							Discount:  0.3,
						},
					},
					TransactionDate: "2023-01-01",
					TotalPrice:      15000,
					BranchId:        "branch-1",
					StaffId:         "staff-1",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{
				repository:   testRepository,
				goodsAdapter: testGoods,
				kafkaAdapter: testKafka,
			}
			if err := c.CreateOfflineOrder(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("CreateOfflineOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
