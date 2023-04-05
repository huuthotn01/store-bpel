package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetGoodsDefault(ctx context.Context, request *schema.GetGoodsDefaultRequest) ([]*schema.GetGoodsDefaultResponseData, error) {
	goods, err := c.repository.GetGoodsDefault(ctx, request.PageSize, request.PageNumber)
	if err != nil {
		return nil, err
	}

	respGoods := make([]*schema.GetGoodsDefaultResponseData, 0, len(goods))

	for _, code := range goods { // code = goodsId
		eachDetail, err := c.getEachProductDetail(ctx, code)
		if err != nil {
			return nil, err
		}

		respGoods = append(respGoods, eachDetail)
	}

	return respGoods, nil
}
