package controller

import (
	"context"
)

func (s *cartServiceController) AddCart(ctx context.Context, request string) error {
	return s.repository.AddCart(ctx, request)
}
