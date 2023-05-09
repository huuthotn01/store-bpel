package controller

import (
	"context"
	"store-bpel/event_service/config"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_UpdateEvent(t *testing.T) {
	type args struct {
		eventId string
		request *schema.UpdateEventRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update event successfully",
			args: args{
				eventId: "event-1",
				request: &schema.UpdateEventRequest{
					Name:      "Test Event",
					Discount:  0.3,
					StartTime: "2023-03-03",
					EndTime:   "2023-03-05",
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
			s := &eventServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := s.UpdateEvent(ctx, tt.args.eventId, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
