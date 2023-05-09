package controller

import (
	"context"
	"reflect"
	"store-bpel/statistic_service/config"
	"store-bpel/statistic_service/schema"
	"testing"
)

func Test_statisticServiceController_GetProfit(t *testing.T) {
	type args struct {
		request *schema.FilterGetStatisticRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetProfitResponseData
		wantErr bool
	}{
		{
			name: "Should get profit correctly",
			args: args{
				request: &schema.FilterGetStatisticRequest{
					BranchId: []string{"branch-1", "branch-2"},
					Gender:   []int{1, 2, 3},
					Start:    "2023-01-01",
					End:      "2023-01-03",
				},
			},
			want: []*schema.GetProfitResponseData{
				{
					Profit: 500,
					Date:   "2023-01-01",
				},
				{
					Profit: 600,
					Date:   "2023-01-02",
				},
				{
					Profit: 10,
					Date:   "2023-01-03",
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
			c := &statisticServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			got, err := c.GetProfit(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProfit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
