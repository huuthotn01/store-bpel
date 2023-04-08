package cart_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/cart_service"
)

func (c *cartBffController) GetCart(ctx context.Context, userId string) (*cart_service.CartData, error) {
	cart, err := c.cartAdapter.GetCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	listGoods := make([]*cart_service.GoodsData, 0, len(cart.Goods))

	for _, good := range cart.Goods {
		listQuantity := make([]*cart_service.QuantityData, 0, len(good.ListQuantity))

		for _, quantity := range good.ListQuantity {
			listQuantity = append(listQuantity, &cart_service.QuantityData{
				GoodsSize:   quantity.GoodsSize,
				GoodsColor:  quantity.GoodsColor,
				Quantity:    quantity.Quantity,
				MaxQuantity: quantity.MaxQuantity,
			})
		}

		listGoods = append(listGoods, &cart_service.GoodsData{
			GoodsId:      good.GoodsId,
			Name:         good.Name,
			UnitPrice:    good.UnitPrice,
			Price:        good.Price,
			Images:       good.Images,
			ListQuantity: listQuantity,
			Discount:     good.Discount,
			GoodsType:    good.GoodsType,
			GoodsGender:  good.GoodsGender,
			GoodsAge:     good.GoodsAge,
			Description:  good.Description,
		})
	}

	return &cart_service.CartData{
		CartId: cart.CartId,
		Goods:  listGoods,
	}, nil
}
