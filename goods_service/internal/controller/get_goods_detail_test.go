package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_GetDetailGoods(t *testing.T) {
	type args struct {
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetGoodsResponseData
		wantErr bool
	}{
		{
			name: "Should get goods detail correctly",
			args: args{
				goodsId: "goods-1",
			},
			want: &schema.GetGoodsResponseData{
				GoodsId:   "goods-1",
				GoodsName: "Goods One",
				Classify: []*schema.GetGoodsResponseData_Classify{
					{
						Size:  "XL",
						Color: "red",
					},
					{
						Size:  "XXL",
						Color: "yellow",
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
			got, err := c.GetDetailGoods(ctx, tt.args.goodsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetailGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDetailGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
