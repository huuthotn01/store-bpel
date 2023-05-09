package controller

import (
	"context"
	"store-bpel/cart_service/config"
	"store-bpel/cart_service/schema"
	"testing"
)

func Test_cartServiceController_AddGoods(t *testing.T) {
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
			name: "Should add goods successfully",
			args: args{
				cartId: "customer-1",
				request: []*schema.AddGoodsRequest{
					{
						GoodsId:    "goods-1",
						GoodsColor: "red",
						GoodsSize:  "XL",
						Quantity:   2,
					},
					{
						GoodsId:    "goods-2",
						GoodsColor: "yellow",
						GoodsSize:  "S",
						Quantity:   5,
					},
				},
			},
		},
	}

	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cartServiceController{
				cfg:          cfg,
				repository:   testRepository,
				goodsAdapter: testGoods,
			}
			if err := s.AddGoods(ctx, tt.args.cartId, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
