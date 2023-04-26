package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

func (b *branchBffController) DeleteBranch(ctx context.Context, request *branch_schema.DeleteBranchRequest) error {
	return b.branchAdapter.DeleteBranch(ctx, request.BranchId)
}
