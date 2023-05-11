package controller

import (
	"context"
	"reflect"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_GetEventCurrent(t *testing.T) {
	type args struct {
		date int
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetEventData
		wantErr bool
	}{
		{
			name: "Should get all current event correctly",
			args: args{
				date: 3,
			},
			want: []*schema.GetEventData{
				{
					Id:       "event-1",
					Name:     "Test Event 1",
					Discount: 0.5,
					Goods:    []string{"goods-1", "goods-2"},
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
			got, err := s.GetEventCurrent(ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventCurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventCurrent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
