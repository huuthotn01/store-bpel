package controller

import (
	"context"
	"store-bpel/event_service/config"
	"testing"
)

func Test_eventServiceController_DeleteEvent(t *testing.T) {
	type args struct {
		eventId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete event successfully",
			args: args{
				eventId: "event-1",
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
			if err := s.DeleteEvent(ctx, tt.args.eventId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
