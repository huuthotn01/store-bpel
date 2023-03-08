package branch_service

import (
	"context"
	branch_schema "store-bpel/bff/admin_bff/schema/branch_service"
)

func (b *branchBffController) GetBranchDetail(ctx context.Context, branchId string) (*branch_schema.GetBranchResponseData, error) {
	branch, err := b.branchAdapter.GetBranchDetail(ctx, branchId)
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

func (b *branchBffController) GetBranch(ctx context.Context) ([]*branch_schema.GetBranchResponseData, error) {
	branch, err := b.branchAdapter.GetBranch(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*branch_schema.GetBranchResponseData, 0, len(branch))
	for _, data := range branch {
		resp = append(resp, &branch_schema.GetBranchResponseData{
			BranchCode:     data.BranchCode,
			BranchName:     data.BranchName,
			BranchProvince: data.BranchProvince,
			BranchDistrict: data.BranchDistrict,
			BranchWard:     data.BranchWard,
			BranchStreet:   data.BranchStreet,
			CreatedAt:      data.CreatedAt,
			Manager:        data.Manager,
			Open:           data.OpenTime,
			Close:          data.CloseTime,
		})
	}

	return resp, nil
}
