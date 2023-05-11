package controller

import (
	"context"
	"store-bpel/cart_service/schema"
	"testing"
)

func Test_cartServiceController_DeleteGoods(t *testing.T) {
	type args struct {
		cartId  string
		request []*schema.DeleteGoodsRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete goods successfully",
			args: args{
				cartId: "customer-1",
				request: []*schema.DeleteGoodsRequest{
					{
						GoodsId:    "goods-1",
						GoodsColor: "red",
						GoodsSize:  "XL",
					},
					{
						GoodsId:    "goods-2",
						GoodsColor: "yellow",
						GoodsSize:  "S",
					},
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cartServiceController{
				repository: testRepository,
			}
			if err := s.DeleteGoods(ctx, tt.args.cartId, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
