package controller

import (
	"context"
	"reflect"
	"store-bpel/cart_service/config"
	"store-bpel/cart_service/schema"
	"testing"
)

func Test_cartServiceController_GetCart(t *testing.T) {
	type args struct {
		request string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.CartData
		wantErr bool
	}{
		{
			name: "Should get cart data correctly",
			args: args{
				request: "customer-1",
			},
			want: &schema.CartData{
				CartId: "customer-1",
				Goods: []*schema.GoodsData{
					{
						GoodsId: "goods-1",
						ListQuantity: []*schema.QuantityData{
							{
								GoodsSize:   "XL",
								GoodsColor:  "red",
								Quantity:    2,
								MaxQuantity: 5,
							},
							{
								GoodsSize:   "S",
								GoodsColor:  "yellow",
								Quantity:    1,
								MaxQuantity: 5,
							},
						},
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
			got, err := s.GetCart(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}
