package controller

import (
	"context"
	"store-bpel/cart_service/repository"
	"store-bpel/cart_service/schema"
)

func (s *cartServiceController) AddGoods(ctx context.Context, cartId int, request []*schema.AddGoodsRequest) error {
	var goodsList []*repository.AddGoodsData
	for _, goods := range request {
		goodsList = append(goodsList, &repository.AddGoodsData{
			GoodsId:    goods.GoodsId,
			GoodsSize:  goods.GoodsSize,
			GoodsColor: goods.GoodsColor,
			Quantity:   goods.Quantity,
		})
	}

	return s.repository.AddGoods(ctx, cartId, goodsList)
}
