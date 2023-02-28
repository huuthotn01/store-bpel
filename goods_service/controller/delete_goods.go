package controller

import "context"

func (c *goodsServiceController) DeleteGoods(ctx context.Context, goodsId string) error {
	return c.repository.UpdateGoodsIsForSaleToNo(ctx, goodsId)
}
