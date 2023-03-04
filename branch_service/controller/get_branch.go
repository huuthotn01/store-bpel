package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"store-bpel/branch_service/schema"
)

func (s *branchServiceController) GetBranch(ctx context.Context) ([]*schema.GetBranchResponseData, error) {
	branches, err := s.repository.GetBranch(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetBranchResponseData, 0, len(branches))
	for _, branch := range branches {
		converted := schema.GetBranchResponseData(*branch)
		res = append(res, &converted)
	}
	return res, nil
}

func (s *branchServiceController) GetBranchDetail(ctx context.Context, branchId string) (*schema.GetBranchResponseData, error) {
	branch, err := s.repository.GetBranchDetail(ctx, branchId)
	if err != nil {
		// no data with given branchId => return nothing
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	converted := schema.GetBranchResponseData(*branch)
	return &converted, nil
}

func (s *branchServiceController) GetBranchStaff(ctx context.Context, branchId string) ([]string, error) {
	res, err := s.repository.GetBranchStaff(ctx, branchId)
	if err != nil {
		return nil, err
	}
	staff := make([]string, 0, len(res))
	for _, s := range res {
		staff = append(staff, s.StaffCode)
	}
	return staff, nil
}
