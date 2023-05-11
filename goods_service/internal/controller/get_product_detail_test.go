package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_GetProductDetail(t *testing.T) {
	type args struct {
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetGoodsDefaultResponseData
		wantErr bool
	}{
		{
			name: "Should get product detail correctly",
			args: args{
				goodsId: "goods-1",
			},
			want: &schema.GetGoodsDefaultResponseData{
				GoodsId: "goods-1",
				Name:    "Goods One",
				Images:  []string{"url-1", "url-2"},
				ListQuantity: []*schema.GetGoodsDefault_QuantityList{
					{
						GoodsColor: "red",
						GoodsSize:  "XL",
						Quantity:   100,
					},
					{
						GoodsColor: "yellow",
						GoodsSize:  "XXL",
						Quantity:   100,
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
			got, err := c.GetProductDetail(ctx, tt.args.goodsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
