package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) SearchGoods(ctx context.Context, request *schema.SearchGoodsRequest) ([]*schema.GetGoodsDefaultResponseData, error) {
	goods, err := c.repository.FilterGoods(ctx, request.Query, request.Category == 2)
	if err != nil {
		return nil, err
	}

	res := make([]*schema.GetGoodsDefaultResponseData, request.PageSize)

	switch request.Category {
	case 0: // Search goods
		for i := 0; i < request.PageSize; i++ {
			goodsDefault, err := c.getEachProductDetail(ctx, goods[i])
			if err != nil {
				return nil, err
			}
			res[i] = goodsDefault
		}
	case 1: // get best-selling
		mapGoodsCodeToTrue := make(map[string]bool, 0)
		for _, goodsId := range goods {
			mapGoodsCodeToTrue[goodsId] = true
		}
		bestGoods, err := c.orderServiceAdapter.GetBestSellingGoods(ctx)
		if err != nil {
			return nil, err
		}
		counter := 0
		for _, goodsId := range bestGoods {
			if _, ok := mapGoodsCodeToTrue[goodsId]; ok {
				goodsDefault, err := c.getEachProductDetail(ctx, goodsId)
				if err != nil {
					return nil, err
				}
				res[counter] = goodsDefault
				counter++
				if counter == request.PageSize {
					break
				}
			}
		}
	case 2: // get newly added
		for i := 0; i < request.PageSize; i++ {
			goodsDefault, err := c.getEachProductDetail(ctx, goods[i])
			if err != nil {
				return nil, err
			}
			res[i] = goodsDefault
		}
	}

	return res, nil
}
