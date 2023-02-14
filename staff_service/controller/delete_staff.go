package controller

import "context"

func (s *staffServiceController) DeleteStaff(ctx context.Context, staffId string) error {
	return s.repository.DeleteStaff(ctx, staffId)
}
