package controller

import (
	"context"
	"store-bpel/branch_service/repository"
	"store-bpel/branch_service/schema"
)

func (s *branchServiceController) UpdateBranch(ctx context.Context, request *schema.UpdateBranchRequest, branchId string) error {
	return s.repository.UpdateBranch(ctx, &repository.BranchModel{
		BranchCode:     branchId,
		BranchName:     request.Name,
		BranchProvince: request.Province,
		BranchDistrict: request.District,
		BranchWard:     request.Ward,
		BranchStreet:   request.Street,
		OpenTime:       request.Open,
		CloseTime:      request.Close,
	})
}

func (s *branchServiceController) UpdateBranchManager(ctx context.Context, request *schema.UpdateBranchManagerRequest, branchId string) error {
	return s.repository.UpdateBranchManager(ctx, branchId, request.StaffId)
}
