package controller

import (
	"context"
	"store-bpel/branch_service/repository"
	"store-bpel/branch_service/schema"
)

func (s *branchServiceController) AddBranchStaff(ctx context.Context, request *schema.AddBranchStaffRequest) error {
	return s.repository.AddBranchStaff(ctx, &repository.BranchStaffModel{
		BranchCode: request.BranchId,
		StaffCode:  request.StaffId,
	})
}
