package controller

import (
	"context"
	"store-bpel/cart_service/schema"
)

func (s *cartServiceController) GetCart(ctx context.Context, request string) (*schema.CartData, error) {
	cartModel, err := s.repository.GetCart(ctx, request)

	if err != nil {
		return nil, err
	}

	goodsList := make([]*schema.GoodsData, 0, len(cartModel.Goods))

	for _, goods := range cartModel.Goods {
		goodsList = append(goodsList, &schema.GoodsData{
			GoodsId:    goods.GoodsId,
			GoodsColor: goods.GoodsColor,
			GoodsSize:  goods.GoodsSize,
			Quantity:   goods.Quantity,
		})
	}

	result := &schema.CartData{
		CartId: cartModel.CartId,
		Goods:  goodsList,
	}
	return result, nil
}
