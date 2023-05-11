package controller

import (
	"context"
	"store-bpel/event_service/schema"
	"testing"
)

func Test_eventServiceController_UploadImage(t *testing.T) {
	type args struct {
		request *schema.UploadImageRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should upload image successfully",
			args: args{
				request: &schema.UploadImageRequest{
					EventId:  "event-1",
					ImageUrl: "image-url",
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
			if err := s.UploadImage(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UploadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
