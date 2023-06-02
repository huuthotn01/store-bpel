package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_handleCustReturn(t *testing.T) {
	type args struct {
		request *schema.CreateGoodsTransactionRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should do customer return successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.handleCustReturn(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("handleCustReturn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
