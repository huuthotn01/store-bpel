package controller

import "context"

func (c *goodsServiceController) DeleteGoodsImage(ctx context.Context, url string) error {
	return c.repository.DeleteGoodsImage(ctx, url)
}
