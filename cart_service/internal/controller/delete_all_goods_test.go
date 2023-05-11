package controller

import (
	"context"
	"testing"
)

func Test_cartServiceController_DeleteAllGoods(t *testing.T) {
	type args struct {
		cartId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete all goods successfully",
			args: args{
				cartId: "customer-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cartServiceController{
				repository: testRepository,
			}
			if err := s.DeleteAllGoods(ctx, tt.args.cartId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllGoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
