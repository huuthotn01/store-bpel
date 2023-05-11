package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_GetGoodsDefault(t *testing.T) {
	type args struct {
		request *schema.GetGoodsDefaultRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetGoodsDefaultResponseData
		wantErr bool
	}{
		{
			name: "Should get goods default correctly",
			args: args{
				request: &schema.GetGoodsDefaultRequest{
					PageSize:   5,
					PageNumber: 5,
				},
			},
			want: []*schema.GetGoodsDefaultResponseData{
				{
					GoodsId: "goods-1",
					Name:    "Goods One",
					Images:  []string{"url-1", "url-2"},
					ListQuantity: []*schema.GetGoodsDefault_QuantityList{
						{
							GoodsSize:  "XL",
							GoodsColor: "red",
							Quantity:   100,
						},
						{
							GoodsSize:  "XXL",
							GoodsColor: "yellow",
							Quantity:   100,
						},
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
			got, err := c.GetGoodsDefault(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoodsDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoodsDefault() got = %v, want %v", got, tt.want)
			}
		})
	}
}
