package controller

import (
	"context"
	"encoding/json"
	"math/rand"
	"store-bpel/library/kafka_lib"
	repository2 "store-bpel/order_service/internal/repository"
	"store-bpel/order_service/schema"
	stat_schema "store-bpel/statistic_service/schema"
)

func (c *orderServiceController) CreateOnlineOrder(ctx context.Context, request *schema.MakeOnlineOrderRequest) error {
	orderPublicCode := c.generateOrderPublicCode()

	var (
		orderGoods    = make([]*repository2.GoodsModel, 0, len(request.GoodsList))
		statGoodsData = make([]*stat_schema.AddOrderDataRequest_GoodsData, 0, len(request.GoodsList))
	)
	for _, data := range request.GoodsList {
		orderGoods = append(orderGoods, &repository2.GoodsModel{
			GoodsCode:  data.GoodsId,
			GoodsSize:  data.Size,
			GoodsColor: data.Color,
			Quantity:   data.Quantity,
			UnitPrice:  data.UnitPrice,
			TotalPrice: data.Price,
			Tax:        data.Tax,
			Image:      data.Image,
			GoodsName:  data.Name,
			Promotion:  data.Discount,
		})

		goodsDetail, err := c.goodsAdapter.GetProductDetail(ctx, data.GoodsId)
		if err != nil {
			return err
		}

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

	err := c.repository.CreateOnlineOrder(ctx, &repository2.OnlineOrdersData{
		PublicOrderCode: orderPublicCode,
		TransactionDate: request.TransactionDate,
		TotalPrice:      request.TotalPrice,
		Goods:           orderGoods,
		OnlineOrder: &repository2.OnlineOrdersModel{
			ExpectedDelivery: request.ExpectedDate,
			ShippingFee:      request.ShipFee,
			CustomerId:       request.CustomerId,
			PaymentMethod:    request.PaymentMethod,
			Street:           request.Address.Street,
			Ward:             request.Address.Ward,
			District:         request.Address.District,
			Province:         request.Address.Province,
			CustomerName:     request.NameReceiver,
			CustomerPhone:    request.PhoneReceiver,
			CustomerEmail:    request.EmailReceiver,
			Status:           0, // initial status
		},
	})
	if err != nil {
		return err
	}

	// call stat service to add order data by kafka
	addOrderDataRequest := &stat_schema.AddOrderDataRequest{
		OrderId:         orderPublicCode,
		TransactionDate: request.TransactionDate,
		ShopCode:        "", // no shop in online order
		GoodsData:       statGoodsData,
	}
	addOrderDataReqByte, err := json.Marshal(addOrderDataRequest)
	if err != nil {
		return err
	}

	return c.kafkaAdapter.Publish(ctx, kafka_lib.STATISTIC_SERVICE_TOPIC, addOrderDataReqByte)
}

func (c *orderServiceController) generateOrderPublicCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := make([]byte, 8) // generate random order code fixed in 8-char size
	for i := range res {
		res[i] = letters[rand.Intn(52)] // 52 is total number of english alphabet letters
	}
	return string(res)
}
