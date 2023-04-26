package controller

import "context"

func (s *branchServiceController) DeleteBranch(ctx context.Context, branchId string) error {
	return s.repository.DeleteBranch(ctx, branchId)
}
