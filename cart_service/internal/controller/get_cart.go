package controller

import (
	"context"
	"store-bpel/cart_service/schema"
	goodsSchema "store-bpel/goods_service/schema"
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

		listQuantity := setListQuantity(ctx, goodsInCart.ListQuantity, productDetail)

		if len(listQuantity) == 0 {
			continue
		}

		goodsList = append(goodsList, &schema.GoodsData{
			GoodsId:      goodsInCart.GoodsId,
			Name:         productDetail.Name,
			UnitPrice:    productDetail.UnitPrice,
			Price:        productDetail.Price,
			Images:       productDetail.Images,
			ListQuantity: listQuantity,
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

func setListQuantity(ctx context.Context, quantityData []*schema.QuantityData, productDetail *goodsSchema.GetGoodsDefaultResponseData) []*schema.QuantityData {
	result := make([]*schema.QuantityData, 0, len(quantityData))
	for _, quantity := range quantityData {
		for _, productQuantity := range productDetail.ListQuantity {
			if quantity.GoodsSize == productQuantity.GoodsSize && quantity.GoodsColor == productQuantity.GoodsColor {
				result = append(result, &schema.QuantityData{
					GoodsSize:   quantity.GoodsSize,
					GoodsColor:  quantity.GoodsColor,
					Quantity:    quantity.Quantity,
					MaxQuantity: productQuantity.Quantity,
				})
				break
			}
		}
	}
	return result
}
