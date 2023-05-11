package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_UploadGoodsImage(t *testing.T) {
	type args struct {
		request *schema.UploadImageRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should upload goods image successfully",
			args: args{
				request: &schema.UploadImageRequest{
					GoodsId:    "goods-1",
					GoodsColor: "red",
					Url:        "image-url",
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
			if err := c.UploadGoodsImage(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UploadGoodsImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
