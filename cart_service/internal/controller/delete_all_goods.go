package controller

import (
	"context"
)

func (s *cartServiceController) DeleteAllGoods(ctx context.Context, cartId string) error {
	return s.repository.DeleteAllGoods(ctx, cartId)
}
