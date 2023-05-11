package controller

import (
	"context"
	"store-bpel/cart_service/schema"
	"testing"
)

func Test_cartServiceController_UpdateGoods(t *testing.T) {
	type args struct {
		cartId  string
		request []*schema.AddGoodsRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update goods correctly",
			args: args{
				cartId: "customer-1",
				request: []*schema.AddGoodsRequest{
					{
						GoodsId:    "goods-1",
						GoodsSize:  "XL",
						GoodsColor: "red",
						Quantity:   2,
					},
					{
						GoodsId:    "goods-1",
						GoodsSize:  "S",
						GoodsColor: "yellow",
						Quantity:   1,
					},
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cartServiceController{
				repository:   testRepository,
				goodsAdapter: testGoods,
			}
			if err := s.UpdateGoods(ctx, tt.args.cartId, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
