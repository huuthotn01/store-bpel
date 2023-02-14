package controller

import "context"

func (s *staffServiceController) DeleteAddRequest(ctx context.Context, staffId string) error {
	return s.repository.DeleteAddRequest(ctx, staffId)
}
