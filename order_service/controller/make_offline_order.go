package controller

import (
	"context"
	"encoding/json"
	"store-bpel/library/kafka_lib"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
	stat_schema "store-bpel/statistic_service/schema"
)

func (c *orderServiceController) CreateOfflineOrder(ctx context.Context, request *schema.MakeOfflineOrderRequest) error {
	orderPublicCode := c.generateOrderPublicCode()

	var (
		orderGoods    = make([]*repository.GoodsModel, 0, len(request.GoodsList))
		statGoodsData = make([]*stat_schema.AddOrderDataRequest_GoodsData, 0, len(request.GoodsList))
	)
	for _, data := range request.GoodsList {
		goodsDetail, err := c.goodsAdapter.GetProductDetail(ctx, data.GoodsId)
		if err != nil {
			return err
		}
		var imgUrl string
		if len(goodsDetail.Images) > 0 {
			imgUrl = goodsDetail.Images[0]
		}

		orderGoods = append(orderGoods, &repository.GoodsModel{
			GoodsCode:  data.GoodsId,
			GoodsSize:  data.Size,
			GoodsColor: data.Color,
			Quantity:   data.Quantity,
			UnitPrice:  data.UnitPrice,
			TotalPrice: data.Price,
			Tax:        data.Tax,
			GoodsName:  goodsDetail.Name,
			Image:      imgUrl,
			Promotion:  data.Discount,
		})

		statGoodsData = append(statGoodsData, &stat_schema.AddOrderDataRequest_GoodsData{
			GoodsId:     data.GoodsId,
			GoodsSize:   data.Size,
			GoodsColor:  data.Color,
			GoodsType:   goodsDetail.GoodsType,
			GoodsGender: goodsDetail.GoodsGender,
			GoodsCost:   goodsDetail.UnitCost,
			UnitPrice:   data.UnitPrice,
			Quantity:    data.Quantity,
		})
	}

	err := c.repository.CreateOfflineOrder(ctx, &repository.OfflineOrdersData{
		PublicOrderCode: orderPublicCode,
		TransactionDate: request.TransactionDate,
		TotalPrice:      request.TotalPrice,
		BranchId:        request.BranchId,
		StaffId:         request.StaffId,
		Goods:           orderGoods,
	})
	if err != nil {
		return err
	}

	// call stat service to add order data by kafka
	addOrderDataRequest := &stat_schema.AddOrderDataRequest{
		OrderId:         orderPublicCode,
		TransactionDate: request.TransactionDate,
		ShopCode:        request.BranchId,
		GoodsData:       statGoodsData,
	}
	addOrderDataReqByte, err := json.Marshal(addOrderDataRequest)
	if err != nil {
		return err
	}

	return c.kafkaAdapter.Publish(ctx, kafka_lib.STATISTIC_SERVICE_TOPIC, addOrderDataReqByte)
}
