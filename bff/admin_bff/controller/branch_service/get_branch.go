package branch_service

import (
	"context"
	"store-bpel/branch_service/schema"
)

func (b *branchBffController) GetBranch(ctx context.Context, branchId string) (*schema.GetBranchDetailResponse, error) {
	branch, err := b.branchAdapter.GetBranch(ctx, branchId)
	if err != nil {
		return nil, err
	}
	return branch, nil
}
