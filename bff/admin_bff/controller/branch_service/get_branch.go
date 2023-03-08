package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

func (b *branchBffController) GetBranch(ctx context.Context, branchId string) (*branch_schema.GetBranchResponseData, error) {
	branch, err := b.branchAdapter.GetBranch(ctx, branchId)
	if err != nil {
		return nil, err
	}

	respData := &branch_schema.GetBranchResponseData{
		BranchCode:     branch.BranchCode,
		BranchName:     branch.BranchName,
		BranchProvince: branch.BranchProvince,
		BranchDistrict: branch.BranchDistrict,
		BranchWard:     branch.BranchWard,
		BranchStreet:   branch.BranchStreet,
		CreatedAt:      branch.CreatedAt,
		Manager:        branch.Manager,
		Open:           branch.OpenTime,
		Close:          branch.CloseTime,
	}

	return respData, nil
}
