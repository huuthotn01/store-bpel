package controller

import (
	"context"
	"testing"
)

func Test_goodsServiceController_DeleteGoodsImage(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete goods image successfully",
			args: args{
				url: "image-url",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.DeleteGoodsImage(ctx, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGoodsImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
