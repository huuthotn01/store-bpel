package controller

import (
	"context"
	"reflect"
	"store-bpel/statistic_service/config"
	"store-bpel/statistic_service/schema"
	"testing"
)

func Test_statisticServiceController_GetOverallStat(t *testing.T) {
	type args struct {
		request *schema.CommonGetStatisticRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetOverallStatisticResponseData
		wantErr bool
	}{
		{
			name: "Should get overall stat correctly",
			args: args{
				request: &schema.CommonGetStatisticRequest{
					Start:    "2023-01-01",
					End:      "2023-01-03",
					BranchId: []string{"branch-1", "branch-2"},
				},
			},
			want: []*schema.GetOverallStatisticResponseData{
				{
					Revenue: 1000,
					Profit:  500,
					Date:    "2023-01-01",
				},
				{
					Revenue: 2000,
					Profit:  600,
					Date:    "2023-01-02",
				},
				{
					Revenue: 100,
					Profit:  10,
					Date:    "2023-01-03",
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
			got, err := c.GetOverallStat(ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOverallStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOverallStat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
