package controller

import (
	"context"
	"reflect"
	"store-bpel/event_service/config"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_GetEventDetail(t *testing.T) {
	type args struct {
		eventId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetEventData
		wantErr bool
	}{
		{
			name: "Should get event detail correctly",
			args: args{
				eventId: "event-2",
			},
			want: &schema.GetEventData{
				Id:        "event-2",
				Name:      "Test Event 2",
				Discount:  0.3,
				StartTime: "2023-06-06",
				EndTime:   "2023-12-12",
				Goods:     []string{"goods-3", "goods-4"},
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
			s := &eventServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			got, err := s.GetEventDetail(ctx, tt.args.eventId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
