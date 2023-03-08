package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

func (b *branchBffController) GetBranchStaff(ctx context.Context, request *branch_schema.GetBranchStaffRequest) (*branch_schema.GetBranchStaffResponseData, error) {
	return nil, nil
}
