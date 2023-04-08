package controller

import (
	"context"
	"store-bpel/cart_service/repository"
	"store-bpel/cart_service/schema"
)

func (s *cartServiceController) DeleteGoods(ctx context.Context, cartId string, request []*schema.DeleteGoodsRequest) error {
	var goodsList []*repository.DeleteGoodsData
	for _, goods := range request {
		goodsList = append(goodsList, &repository.DeleteGoodsData{
			GoodsId:    goods.GoodsId,
			GoodsSize:  goods.GoodsSize,
			GoodsColor: goods.GoodsColor,
		})
	}

	return s.repository.DeleteGoods(ctx, cartId, goodsList)
}
