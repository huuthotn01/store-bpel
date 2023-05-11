package controller

import (
	"context"
	"store-bpel/branch_service/schema"
	"testing"
)

func Test_branchServiceController_UpdateBranch(t *testing.T) {
	type args struct {
		request  *schema.UpdateBranchRequest
		branchId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update branch successfully",
			args: args{
				request: &schema.UpdateBranchRequest{
					Name:     "Branch Number One",
					Street:   "To Hien Thanh",
					Ward:     "Phuong 11",
					District: "Quan 10",
					Province: "Thanh pho Ho Chi Minh",
					Open:     "08:30",
					Close:    "17:30",
				},
				branchId: "branch-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			if err := s.UpdateBranch(ctx, tt.args.request, tt.args.branchId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_branchServiceController_UpdateBranchManager(t *testing.T) {
	type args struct {
		request  *schema.UpdateBranchManagerRequest
		branchId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update branch manager successfully",
			args: args{
				request: &schema.UpdateBranchManagerRequest{
					StaffId: "staff-10",
				},
				branchId: "branch-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			if err := s.UpdateBranchManager(ctx, tt.args.request, tt.args.branchId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBranchManager() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
