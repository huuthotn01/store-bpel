package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

func (b *branchBffController) GetBranchStaff(ctx context.Context, request *branch_schema.GetBranchStaffRequest) (*branch_schema.GetBranchStaffResponseData, error) {
	staff, err := b.branchAdapter.GetBranchStaff(ctx, request.BranchId)
	if err != nil {
		return nil, err
	}

	return &branch_schema.GetBranchStaffResponseData{
		Staffs: staff,
	}, nil
}
