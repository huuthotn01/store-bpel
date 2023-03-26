package controller

import (
	"context"
	"math/rand"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
	"time"
)

func (c *orderServiceController) CreateOnlineOrder(ctx context.Context, request *schema.MakeOnlineOrderRequest) error {
	orderPublicCode := c.generateOrderPublicCode()
	expectedDelivery, err := c.getExpectedDelivery(request.TransactionDate)
	if err != nil {
		return err
	}

	orderGoods := make([]*repository.GoodsModel, 0, len(request.GoodsList))
	for _, data := range request.GoodsList {
		orderGoods = append(orderGoods, &repository.GoodsModel{
			GoodsCode:  data.GoodsId,
			GoodsSize:  data.Size,
			GoodsColor: data.Color,
			Quantity:   data.Quantity,
			UnitPrice:  data.UnitPrice,
			TotalPrice: data.Price,
			Tax:        data.Tax,
			Image:      "", // TODO add image link
			Promotion:  data.Discount,
		})
	}

	return c.repository.CreateOnlineOrder(ctx, &repository.OnlineOrdersData{
		PublicOrderCode: orderPublicCode,
		TransactionDate: request.TransactionDate,
		TotalPrice:      request.TotalPrice,
		Goods:           orderGoods,
		OnlineOrder: &repository.OnlineOrdersModel{
			ExpectedDelivery: expectedDelivery,
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
}

func (c *orderServiceController) generateOrderPublicCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := make([]byte, 8) // generate random order code fixed in 8-char size
	for i := range res {
		res[i] = letters[rand.Intn(52)] // 52 is total number of english alphabet letters
	}
	return string(res)
}

func (c *orderServiceController) getExpectedDelivery(transactionDate string) (string, error) {
	createdDate, err := time.Parse("2006-01-02", transactionDate)
	if err != nil {
		return "", err
	}

	// after 5 days
	return createdDate.Add(5 * 24 * time.Hour).Format("2006-01-02"), nil
}
