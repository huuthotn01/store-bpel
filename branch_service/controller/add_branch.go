package controller

import (
	"context"
	"store-bpel/branch_service/repository"
	"store-bpel/branch_service/schema"
)

func (s *branchServiceController) AddBranch(ctx context.Context, request *schema.AddBranchRequest) error {
	branchModel := &repository.BranchModel{
		BranchName:     request.Name,
		BranchProvince: request.Province,
		BranchDistrict: request.District,
		BranchWard:     request.Ward,
		BranchStreet:   request.Street,
		OpenTime:       request.Open,
		CloseTime:      request.Close,
	}
	return s.repository.AddBranch(ctx, branchModel)
}
