package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
	"store-bpel/branch_service/schema"
)

func (b *branchBffController) UpdateBranch(ctx context.Context, request *branch_schema.UpdateBranchRequest) error {
	return b.branchAdapter.UpdateBranch(ctx, request.BranchId, &schema.UpdateBranchRequest{
		Name:     request.Name,
		Street:   request.Street,
		Ward:     request.Ward,
		District: request.District,
		Province: request.Province,
		Open:     request.Open,
		Close:    request.Close,
	})
}
