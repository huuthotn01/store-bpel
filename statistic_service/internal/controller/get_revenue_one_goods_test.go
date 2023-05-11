package controller

import (
	"context"
	"reflect"
	"store-bpel/statistic_service/schema"
	"testing"
)

func Test_statisticServiceController_GetRevenueOneGoods(t *testing.T) {
	type args struct {
		request *schema.CommonGetStatisticRequest
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetRevenueResponseData
		wantErr bool
	}{
		{
			name: "Should get revenue one goods correctly",
			args: args{
				request: &schema.CommonGetStatisticRequest{
					Start:    "2023-01-01",
					End:      "2023-01-03",
					BranchId: []string{"branch-1", "branch-2"},
				},
				goodsId: "goods-1",
			},
			want: []*schema.GetRevenueResponseData{
				{
					Revenue: 1000,
					Date:    "2023-01-01",
				},
				{
					Revenue: 2000,
					Date:    "2023-01-02",
				},
				{
					Revenue: 100,
					Date:    "2023-01-03",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &statisticServiceController{
				repository: testRepository,
			}
			got, err := c.GetRevenueOneGoods(ctx, tt.args.request, tt.args.goodsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRevenueOneGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRevenueOneGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
