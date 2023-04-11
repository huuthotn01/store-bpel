package goods_service

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) UploadImage(ctx context.Context, goodsId, goodsColor, url string, isDefault bool) error {
	return c.goodsAdapter.UploadImage(ctx, &schema.UploadImageRequest{
		GoodsId:    goodsId,
		GoodsColor: goodsColor,
		Url:        url,
		IsDefault:  isDefault,
	})
}
