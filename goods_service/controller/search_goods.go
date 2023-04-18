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
	if len(goods) == 0 {
		return nil, nil
	}

	res := make([]*schema.GetGoodsDefaultResponseData, 0, request.PageSize)

	switch request.Category {
	case 0, 2: // 0: default search goods by filtering name, 2: newly-added
		counter := 0
		for _, data := range goods {
			goodsDefault, err := c.getEachProductDetail(ctx, data)
			if err != nil {
				return nil, err
			}
			res = append(res, goodsDefault)
			counter++
			if counter == request.PageSize {
				break
			}
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
				res = append(res, goodsDefault)
				counter++
				if counter == request.PageSize {
					break
				}
			}
		}
	}

	return res, nil
}
