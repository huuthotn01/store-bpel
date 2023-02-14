package controller

import "context"

func (s *branchServiceController) DeleteBranch(ctx context.Context, branchId int32) error {
	return s.repository.DeleteBranch(ctx, branchId)
}
