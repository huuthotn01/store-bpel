package controller

import "context"

func (s *staffServiceController) DeleteStaff(ctx context.Context, staffId string) error {
	// todo call account service to deactivate account
	return s.repository.DeleteStaff(ctx, staffId)
}
