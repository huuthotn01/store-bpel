package controller

import (
	"context"
	"store-bpel/cart_service/schema"
)

type GoodsClassify struct {
	GoodsId      string
	ListQuantity []*schema.QuantityData
}

func (s *cartServiceController) GetCart(ctx context.Context, request string) (*schema.CartData, error) {
	cart, err := s.repository.GetCart(ctx, request)

	if err != nil {
		return nil, err
	}

	goodsClassify := make([]*GoodsClassify, 0)
	for _, goods := range cart.Goods {
		// check goodsId exists in list classify
		var foundGoods *GoodsClassify
		for _, classified := range goodsClassify {
			if classified.GoodsId == goods.GoodsId {
				foundGoods = classified
				break
			}
		}

		// add goods to list classify if it doesn't exist
		if foundGoods == nil {
			newGoods := &GoodsClassify{
				GoodsId:      goods.GoodsId,
				ListQuantity: []*schema.QuantityData{},
			}
			goodsClassify = append(goodsClassify, newGoods)
			foundGoods = newGoods
		}

		// add QuantityData
		quantityData := &schema.QuantityData{
			GoodsSize:  goods.GoodsSize,
			GoodsColor: goods.GoodsColor,
			Quantity:   goods.Quantity,
		}
		foundGoods.ListQuantity = append(foundGoods.ListQuantity, quantityData)
	}

	goodsList := make([]*schema.GoodsData, 0, len(cart.Goods))
	for _, goodsInCart := range goodsClassify {
		productDetail, err := s.goodsAdapter.GetProductDetail(ctx, goodsInCart.GoodsId)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, &schema.GoodsData{
			GoodsId:      goodsInCart.GoodsId,
			Name:         productDetail.Name,
			UnitPrice:    productDetail.UnitPrice,
			Price:        productDetail.Price,
			Images:       productDetail.Images,
			ListQuantity: goodsInCart.ListQuantity,
			Discount:     productDetail.Discount,
			GoodsType:    productDetail.GoodsType,
			GoodsGender:  productDetail.GoodsGender,
			GoodsAge:     productDetail.GoodsAge,
			Description:  productDetail.Description,
		})

	}

	result := &schema.CartData{
		CartId: cart.CartId,
		Goods:  goodsList,
	}
	return result, nil
}
