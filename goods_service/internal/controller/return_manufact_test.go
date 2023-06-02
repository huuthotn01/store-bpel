package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_handleReturnManufact(t *testing.T) {
	type args struct {
		request *schema.CreateGoodsTransactionRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should return error when db get goods in wh data fail",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode: "invalid-goods-code",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when db get goods in wh data return nothing",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode: "no-data-goods",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when request quantity exceed db quantity",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode: "goods-1",
					Quantity:  10,
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.handleReturnManufact(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("handleReturnManufact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
