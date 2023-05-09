package controller

import (
	"context"
	"store-bpel/event_service/config"
	"testing"
)

func Test_eventServiceController_DeleteImage(t *testing.T) {
	type args struct {
		eventId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete images successfully",
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
			if err := s.DeleteImage(ctx, tt.args.eventId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
