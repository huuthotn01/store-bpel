package controller

import (
	"context"
	"errors"
	"store-bpel/cart_service/internal/repository"
	"store-bpel/cart_service/schema"
)

func (s *cartServiceController) AddGoods(ctx context.Context, cartId string, request []*schema.AddGoodsRequest) error {
	var goodsList []*repository.AddGoodsData
	for _, goods := range request {
		maxQuantity, err := s.getMaxQuantity(ctx, goods)
		if err != nil {
			return err
		}

		if goods.Quantity > maxQuantity {
			return errors.New("AddGoods-quantity limit exceeded")
		}

		goodsList = append(goodsList, &repository.AddGoodsData{
			GoodsId:     goods.GoodsId,
			GoodsSize:   goods.GoodsSize,
			GoodsColor:  goods.GoodsColor,
			Quantity:    goods.Quantity,
			MaxQuantity: maxQuantity,
		})
	}

	return s.repository.AddGoods(ctx, cartId, goodsList)
}

func (s *cartServiceController) getMaxQuantity(ctx context.Context, goodsClassify *schema.AddGoodsRequest) (int, error) {
	productDetail, err := s.goodsAdapter.GetProductDetail(ctx, goodsClassify.GoodsId)
	if err != nil {
		return 0, err
	}

	for _, productQuantity := range productDetail.ListQuantity {
		if goodsClassify.GoodsSize == productQuantity.GoodsSize && goodsClassify.GoodsColor == productQuantity.GoodsColor {
			return productQuantity.Quantity, nil
		}
	}
	return 0, errors.New("getMaxQuantity - Not found goods goodsId = " + goodsClassify.GoodsId + ", goodsColor = " + goodsClassify.GoodsColor + ", goodsSize = " + goodsClassify.GoodsSize)
}
