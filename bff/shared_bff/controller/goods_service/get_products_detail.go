package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/goods_service"
)

func (c *goodsBffController) GetProductsDetail(ctx context.Context, request *goods_service.GetProductsDetailRequest) (*goods_service.GetGoodsDefaultResponseData, error) {
	product, err := c.goodsAdapter.GetProductDetail(ctx, request.GoodsId)
	if err != nil {
		return nil, err
	}

	quantityList := make([]*goods_service.GetGoodsDefault_QuantityList, 0, len(product.ListQuantity))
	for _, ele := range product.ListQuantity {
		quantityList = append(quantityList, &goods_service.GetGoodsDefault_QuantityList{
			GoodsSize:  ele.GoodsSize,
			GoodsColor: ele.GoodsColor,
			Quantity:   ele.Quantity,
		})
	}

	return &goods_service.GetGoodsDefaultResponseData{
		GoodsId:      product.GoodsId,
		Name:         product.Name,
		UnitPrice:    product.UnitPrice,
		Price:        product.Price,
		Images:       append([]string{}, product.Images...),
		ListQuantity: quantityList,
		Discount:     product.Discount,
		GoodsType:    product.GoodsType,
		GoodsGender:  product.GoodsGender,
		GoodsAge:     product.GoodsAge,
	}, nil
}
