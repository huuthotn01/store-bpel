package controller

import (
	"context"
	"reflect"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_CheckWarehouse(t *testing.T) {
	type args struct {
		request *schema.CheckWarehouseRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.CheckWarehouseResponseData
		wantErr bool
	}{
		{
			name: "Should return correct actions",
			args: args{
				request: &schema.CheckWarehouseRequest{
					Elements: []*schema.CheckWarehouseRequestElement{
						{
							GoodsCode:  "goods-1",
							GoodsSize:  "XL",
							GoodsColor: "red",
							Quantity:   10,
						},
					},
				},
			},
			want: &schema.CheckWarehouseResponseData{
				NeedTransfer: true,
				WarehouseActions: []*schema.WarehouseActions{
					{
						GoodsCode:  "goods-1",
						GoodsSize:  "XL",
						GoodsColor: "red",
						Quantity:   3,
						From:       "warehouse-1",
						To:         "warehouse-2",
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
			got, err := c.CheckWarehouse(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWarehouse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
