package controller

import (
	"context"
	"store-bpel/statistic_service/schema"
	"testing"
)

func Test_statisticServiceController_AddOrderData(t *testing.T) {
	type args struct {
		request *schema.AddOrderDataRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add order data successfully",
			args: args{
				request: &schema.AddOrderDataRequest{
					OrderId:         "order-id-1",
					ShopCode:        "shop-1",
					TransactionDate: "2023-02-02",
					GoodsData: []*schema.AddOrderDataRequest_GoodsData{
						{
							GoodsId:     "goods-1",
							GoodsSize:   "XL",
							GoodsColor:  "red",
							GoodsType:   "none",
							GoodsGender: 1,
							GoodsCost:   1000,
							UnitPrice:   1500,
							Quantity:    1,
						},
						{
							GoodsId:     "goods-2",
							GoodsSize:   "L",
							GoodsColor:  "yellow",
							GoodsType:   "none",
							GoodsGender: 2,
							GoodsCost:   100,
							UnitPrice:   50,
							Quantity:    3,
						},
					},
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &statisticServiceController{
				repository: testRepository,
			}
			if err := c.AddOrderData(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddOrderData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
