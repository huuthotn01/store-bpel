package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_GetWarehouseByGoods(t *testing.T) {
	type args struct {
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetGoodsInWarehouseResponseData
		wantErr bool
	}{
		{
			name: "Should get warehouse by goods successfully",
			args: args{
				goodsId: "goods-1",
			},
			want: []*schema.GetGoodsInWarehouseResponseData{
				{
					GoodsCode:  "goods-1",
					GoodsSize:  "XXL",
					GoodsColor: "red",
					WhCode:     "wh-1",
					Quantity:   2,
				},
			},
		},
		{
			name: "Should return error when db get warehouse by goods fail",
			args: args{
				goodsId: "invalid-goods-wh",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			got, err := c.GetWarehouseByGoods(ctx, tt.args.goodsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWarehouseByGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWarehouseByGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
