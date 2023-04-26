package controller

import "context"

func (s *staffServiceController) DeleteStaff(ctx context.Context, staffId string) error {
	// todo call account service to deactivate account
	return s.repository.DeleteStaffUpdateStatus(ctx, staffId)
}
