package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
	"store-bpel/branch_service/schema"
)

func (b *branchBffController) UpdateBranchManager(ctx context.Context, request *branch_schema.UpdateBranchManagerRequest) error {
	return b.branchAdapter.UpdateBranchManager(ctx, request.BranchId, &schema.UpdateBranchManagerRequest{
		StaffId: request.StaffId,
	})
}
