package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_UpdateGoods(t *testing.T) {
	type args struct {
		request []*schema.UpdateGoodsRequest
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update goods successfully",
			args: args{
				request: []*schema.UpdateGoodsRequest{
					{
						GoodsSize:  "XL",
						GoodsColor: "red",
						UnitPrice:  2000,
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
			if err := c.UpdateGoods(ctx, tt.args.request, tt.args.goodsId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
