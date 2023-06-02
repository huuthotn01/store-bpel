package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_SearchGoods(t *testing.T) {
	type args struct {
		request *schema.SearchGoodsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetGoodsDefaultResponseData
		wantErr bool
	}{
		{
			name: "Should return result correctly, case default",
			args: args{
				request: &schema.SearchGoodsRequest{
					Query:    "goods-1",
					PageSize: 5,
					Category: 0,
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
		{
			name: "Should return result correctly, case best-selling",
			args: args{
				request: &schema.SearchGoodsRequest{
					Query:    "goods-1",
					PageSize: 5,
					Category: 1,
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
		{
			name: "Should return error when db filter goods fail",
			args: args{
				request: &schema.SearchGoodsRequest{
					Query:    "invalid-name",
					PageSize: 5,
					Category: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "Should return nothing when db filter goods return nothing",
			args: args{
				request: &schema.SearchGoodsRequest{
					Query:    "empty-name",
					PageSize: 5,
					Category: 1,
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository:          testRepository,
				orderServiceAdapter: testOrder,
			}
			got, err := c.SearchGoods(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
