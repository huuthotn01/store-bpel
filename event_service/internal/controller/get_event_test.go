package controller

import (
	"context"
	"reflect"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_GetEvent(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		ctx     context.Context
		want    []*schema.GetEventData
		wantErr bool
	}{
		{
			name: "Should get all events correctly",
			ctx:  ctx,
			want: []*schema.GetEventData{
				{
					Id:        "event-1",
					Name:      "Test Event 1",
					Discount:  0.5,
					StartTime: "2023-01-01",
					EndTime:   "2023-06-01",
					Goods:     []string{"goods-1", "goods-2"},
				},
				{
					Id:        "event-2",
					Name:      "Test Event 2",
					Discount:  0.3,
					StartTime: "2023-01-01",
					EndTime:   "2023-01-05",
					Goods:     []string{"goods-3", "goods-4"},
				},
			},
		},
		{
			name:    "Should return error when event not valid",
			ctx:     context.WithValue(ctx, "status", "invalid-event"),
			wantErr: true,
		},
		{
			name:    "Should return error when db return error get event",
			ctx:     context.WithValue(ctx, "status", "db-fail"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &eventServiceController{
				repository: testRepository,
			}
			got, err := s.GetEvent(tt.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
