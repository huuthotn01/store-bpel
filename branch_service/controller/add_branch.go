package controller

import (
	"context"
	"fmt"
	"store-bpel/branch_service/repository"
	"store-bpel/branch_service/schema"
	"time"
)

func (s *branchServiceController) AddBranch(ctx context.Context, request *schema.AddBranchRequest) error {
	branchId := fmt.Sprintf("branch_%d", time.Now().Unix())

	branchModel := &repository.BranchModel{
		BranchCode:     branchId,
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
