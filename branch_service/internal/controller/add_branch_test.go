package controller

import (
	"context"
	"store-bpel/branch_service/schema"
	"testing"
)

func Test_branchServiceController_AddBranch(t *testing.T) {
	type args struct {
		request *schema.AddBranchRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add branch successfully",
			args: args{
				request: &schema.AddBranchRequest{
					Name:     "Branch One",
					Street:   "Ly Thuong Kiet",
					Ward:     "Ward 11",
					District: "District 10",
					Province: "HCMC",
					Open:     "07:00",
					Close:    "21:00",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			if err := s.AddBranch(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
