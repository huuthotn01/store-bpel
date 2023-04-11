package goods_service

import (
	"context"
	"store-bpel/bff/shared_bff/schema/goods_service"
	"store-bpel/goods_service/schema"
)

func (c *goodsBffController) GetGoodsDefault(ctx context.Context, request *goods_service.GetGoodsDefaultRequest) ([]*goods_service.GetGoodsDefaultResponseData, error) {
	goods, err := c.goodsAdapter.GetGoodsDefault(ctx, &schema.GetGoodsDefaultRequest{
		PageNumber: request.PageNumber,
		PageSize:   request.PageSize,
	})
	if err != nil {
		return nil, err
	}

	respGoods := make([]*goods_service.GetGoodsDefaultResponseData, 0, len(goods))
	for _, data := range goods {
		quantityList := make([]*goods_service.GetGoodsDefault_QuantityList, 0, len(data.ListQuantity))
		for _, ele := range data.ListQuantity {
			quantityList = append(quantityList, &goods_service.GetGoodsDefault_QuantityList{
				GoodsSize:  ele.GoodsSize,
				GoodsColor: ele.GoodsColor,
				Quantity:   ele.Quantity,
			})
		}
		respGoods = append(respGoods, &goods_service.GetGoodsDefaultResponseData{
			GoodsId:   data.GoodsId,
			Name:      data.Name,
			UnitPrice: data.UnitPrice,
			// Price:        data.Price,
			Price:        data.UnitPrice,
			Images:       append([]string{}, data.Images...),
			ListQuantity: quantityList,
			// Discount:     data.Discount,
			Discount:    0,
			GoodsType:   data.GoodsType,
			GoodsGender: data.GoodsGender,
			GoodsAge:    data.GoodsAge,
		})
	}

	return respGoods, nil
}
