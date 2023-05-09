package controller

import (
	"context"
	"store-bpel/branch_service/config"
	"testing"
)

func Test_branchServiceController_DeleteBranch(t *testing.T) {
	type args struct {
		branchId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete branch successfully",
			args: args{
				branchId: "branch-1",
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
			s := &branchServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := s.DeleteBranch(ctx, tt.args.branchId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
