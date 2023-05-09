package controller

import (
	"context"
	"store-bpel/customer_service/config"
	"testing"
)

func Test_customerServiceController_DeleteImage(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete image successfully",
			args: args{
				username: "httn",
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
			c := &customerServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := c.DeleteImage(ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
