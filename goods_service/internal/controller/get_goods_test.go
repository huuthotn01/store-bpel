package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_GetGoods(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetGoodsResponseData
		wantErr bool
	}{
		{
			name: "Should get goods correctly",
			want: []*schema.GetGoodsResponseData{
				{
					GoodsId:   "goods-1",
					GoodsName: "Goods One",
					Classify: []*schema.GetGoodsResponseData_Classify{
						{
							Color: "red",
							Size:  "XL",
						},
						{
							Color: "yellow",
							Size:  "XXL",
						},
					},
					Image: []string{"url-1", "url-2"},
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
			got, err := c.GetGoods(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
