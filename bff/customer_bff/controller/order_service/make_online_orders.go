package order_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/order_service"
	"store-bpel/order_service/schema"
)

func (c *orderBffController) CreateOnlineOrder(ctx context.Context, request *order_service.MakeOnlineOrderRequest) error {
	coreGoods := make([]*schema.OrderGoodsRequest, 0, len(request.GoodsList))

	for _, data := range request.GoodsList {
		coreGoods = append(coreGoods, &schema.OrderGoodsRequest{
			GoodsId:   data.GoodsId,
			UnitPrice: data.UnitPrice,
			Price:     data.Price,
			Quantity:  data.Quantity,
			Size:      data.Size,
			Color:     data.Color,
			Discount:  data.Discount,
			Tax:       data.Tax,
		})
	}

	return c.orderAdapter.CreateOnlineOrders(ctx, &schema.MakeOnlineOrderRequest{
		CustomerId:      request.CustomerId,
		PaymentMethod:   request.PaymentMethod,
		TotalPrice:      request.TotalPrice,
		ShipFee:         request.ShipFee,
		TransactionDate: request.TransactionDate,
		ExpectedDate:    request.ExpectedDate,
		NameReceiver:    request.NameReceiver,
		PhoneReceiver:   request.PhoneReceiver,
		EmailReceiver:   request.EmailReceiver,
		GoodsList:       coreGoods,
		Address: &schema.Address{
			Street:   request.Address.Street,
			Ward:     request.Address.Ward,
			District: request.Address.District,
			Province: request.Address.Province,
		},
	})
}
