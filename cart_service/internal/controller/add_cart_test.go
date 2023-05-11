package controller

import (
	"context"
	"testing"
)

func Test_cartServiceController_AddCart(t *testing.T) {
	type args struct {
		request string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add cart successfully",
			args: args{
				request: "customer-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cartServiceController{
				repository: testRepository,
			}
			if err := s.AddCart(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
