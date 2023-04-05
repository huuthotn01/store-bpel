package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) GetProductDetail(ctx context.Context, goodsId string) ([]*schema.GetGoodsDefaultResponseData, error) {
	detail, err := c.getEachProductDetail(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	return []*schema.GetGoodsDefaultResponseData{detail}, nil
}

func (c *goodsServiceController) getEachProductDetail(ctx context.Context, goodsId string) (*schema.GetGoodsDefaultResponseData, error) {
	goodsDetail, err := c.repository.GetDetailGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	listQuantity := make([]*schema.GetGoodsDefault_QuantityList, 0, len(goodsDetail))
	for _, g := range goodsDetail {
		listQuantity = append(listQuantity, &schema.GetGoodsDefault_QuantityList{
			GoodsSize:  g.GoodsSize,
			GoodsColor: g.GoodsColor,
			Quantity:   100, // TODO call warehouse to get total goods quantity in all warehouses
		})
	}

	images, err := c.repository.GetGoodsImages(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	imgList := make([]string, 0, len(images))
	for _, data := range images {
		imgList = append(imgList, data.GoodsImg)
	}

	// get promotion
	promo, err := c.eventServiceAdapter.GetEventByGoods(ctx, goodsId)
	if err != nil {
		return nil, err
	}

	salePrice := 0
	if goodsDetail[0].UnitPrice-int(promo[0].Discount) > 0 {
		salePrice = goodsDetail[0].UnitPrice - int(promo[0].Discount)
	}

	return &schema.GetGoodsDefaultResponseData{
		GoodsId:      goodsId,
		Name:         goodsDetail[0].GoodsName,
		UnitPrice:    goodsDetail[0].UnitPrice,
		Discount:     int(promo[0].Discount),
		Price:        salePrice,
		GoodsType:    goodsDetail[0].GoodsType,
		GoodsAge:     goodsDetail[0].GoodsAge,
		GoodsGender:  goodsDetail[0].GoodsGender,
		Images:       imgList,
		ListQuantity: listQuantity,
		Description:  goodsDetail[0].Description,
	}, nil
}
