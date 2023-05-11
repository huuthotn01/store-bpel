package controller

import (
	"context"
	"testing"
)

func Test_goodsServiceController_DeleteGoods(t *testing.T) {
	type args struct {
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete goods correctly",
			args: args{
				goodsId: "goods-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.DeleteGoods(ctx, tt.args.goodsId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
