package controller

import (
	"context"
	"store-bpel/customer_service/schema"
	"testing"
)

func Test_customerServiceController_UploadImage(t *testing.T) {
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
					Username: "httn",
					ImageUrl: "image-url",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customerServiceController{
				repository: testRepository,
			}
			if err := c.UploadImage(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UploadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
