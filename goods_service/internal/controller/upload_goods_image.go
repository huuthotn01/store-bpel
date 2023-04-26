package controller

import (
	"context"
	"store-bpel/goods_service/internal/repository"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) UploadGoodsImage(ctx context.Context, request *schema.UploadImageRequest) error {
	return c.repository.AddGoodsImage(ctx, &repository.GoodsImg{
		GoodsCode:  request.GoodsId,
		GoodsColor: request.GoodsColor,
		GoodsImg:   request.Url,
		IsDefault:  request.IsDefault,
	})
}
