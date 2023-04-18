package goods_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/goods_service"
)

func (c *goodsBffController) DeleteImage(ctx context.Context, request *goods_service.DeleteImageRequest) error {
	return c.goodsAdapter.DeleteImage(ctx, request.Url)
}
