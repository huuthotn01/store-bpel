package controller

import (
	"context"
	"reflect"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_GetEventByGoods(t *testing.T) {
	type args struct {
		goodsId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetEventByGoodsData
		wantErr bool
	}{
		{
			name: "Should get event by goods correctly",
			args: args{
				goodsId: "goods-1",
			},
			want: []*schema.GetEventByGoodsData{
				{
					Id:       "event-1",
					Name:     "Test Event 1",
					Discount: 0.5,
				},
				{
					Id:       "event-2",
					Name:     "Test Event 2",
					Discount: 0.3,
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &eventServiceController{
				repository: testRepository,
			}
			got, err := s.GetEventByGoods(ctx, tt.args.goodsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventByGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventByGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
