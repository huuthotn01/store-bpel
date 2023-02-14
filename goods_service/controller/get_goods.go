package controller

import (
	"context"
	"store-bpel/goods_service/schema"
)

func (s *goodsServiceController) GetGoods(ctx context.Context) (*schema.GetGoodsResponse, error) {
	goods, err := s.repository.GetGoods(ctx)
	res := make([]*schema.GoodsModel, 0, len(goods))
	for _, item := range goods {
		converted := schema.GoodsModel(*item)
		res = append(res, &converted)
	}
	if err != nil {
		return nil, err
	}
	_, err = s.warehouseServiceAdapter.GetWarehouse(ctx)
	if err != nil {
		return nil, err
	}
	return &schema.GetGoodsResponse{
		StatusCode: 200,
		Message: "OK",
		Result: res,
	}, nil
}