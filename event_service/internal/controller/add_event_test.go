package controller

import (
	"context"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_AddEvent(t *testing.T) {
	type args struct {
		request *schema.AddEventRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add event successfully",
			args: args{
				request: &schema.AddEventRequest{
					Name:      "Test Discount",
					Discount:  0.5,
					StartTime: "2023-01-01",
					EndTime:   "2023-12-31",
					Goods:     []string{"goods-1", "goods-2"},
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
			if err := s.AddEvent(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
