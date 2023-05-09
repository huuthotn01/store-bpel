package controller

import (
	"context"
	"store-bpel/branch_service/config"
	"store-bpel/branch_service/schema"
	"testing"
)

func Test_branchServiceController_AddBranchStaff(t *testing.T) {
	type args struct {
		request *schema.AddBranchStaffRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add branch staff successfully",
			args: args{
				request: &schema.AddBranchStaffRequest{
					StaffId:  "staff-1",
					BranchId: "branch-1",
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
			s := &branchServiceController{
				cfg:        cfg,
				repository: testRepository,
			}
			if err := s.AddBranchStaff(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddBranchStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
