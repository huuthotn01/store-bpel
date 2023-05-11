package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_AddGoods(t *testing.T) {
	type args struct {
		request []*schema.AddGoodsRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add goods successfully",
			args: args{
				request: []*schema.AddGoodsRequest{
					{
						GoodsSize:  "XL",
						GoodsColor: "red",
						UnitPrice:  5000,
						UnitCost:   1000,
					},
					{
						GoodsSize:  "XXL",
						GoodsColor: "dark",
						UnitPrice:  5000,
						UnitCost:   1000,
					},
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.AddGoods(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
