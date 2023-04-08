package controller

import (
	"context"
	"errors"
	"store-bpel/cart_service/repository"
	"store-bpel/cart_service/schema"
)

func (s *cartServiceController) UpdateGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error {
	var goodsList []*repository.AddGoodsData
	for _, goods := range request {
		maxQuantity, err := s.getMaxQuantity(ctx, goods)
		if err != nil {
			return err
		}

		if goods.Quantity > maxQuantity {
			return errors.New("UpdateGoods-quantity limit exceeded")
		}
		goodsList = append(goodsList, &repository.AddGoodsData{
			GoodsId:    goods.GoodsId,
			GoodsSize:  goods.GoodsSize,
			GoodsColor: goods.GoodsColor,
			Quantity:   goods.Quantity,
		})
	}

	return s.repository.UpdateGoods(ctx, cartId, goodsList)
}
