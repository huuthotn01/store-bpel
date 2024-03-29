package controller

import (
	"context"
	"store-bpel/order_service/internal/repository"
	"store-bpel/order_service/schema"
)

type OrderDetailMappingData struct {
	ListGoods     []*schema.OrderGoodsResponse
	GoodsNum      int
	TotalDiscount int
	StatusShip    []*schema.GetListOrderStateResponse
}

func (c *orderServiceController) GetOrderDetail(ctx context.Context, orderId string) (*schema.GetOrderDetailCustomerResponseData, error) {
	privateOrderId, err := c.repository.GetPrivateOrderCode(ctx, orderId)
	if err != nil {
		return nil, err
	}

	order, err := c.repository.GetOnlineOrderDetail(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	mappingData := c.mapOrderDetailData(order)

	return &schema.GetOrderDetailCustomerResponseData{
		// OrderId is private information => not return for customer
		OrderCode:       order.OrderData.PublicOrderCode,
		PaymentMethod:   order.OnlineOrderData.PaymentMethod,
		ListGoods:       mappingData.ListGoods,
		TotalPrice:      order.OrderData.TotalPrice,
		TotalGoods:      mappingData.GoodsNum,
		TotalDiscount:   mappingData.TotalDiscount,
		TotalOrder:      order.OrderData.TotalPrice + order.OnlineOrderData.ShippingFee,
		IsCompleted:     order.OnlineOrderData.Status == 4,
		ShipFee:         order.OnlineOrderData.ShippingFee,
		StatusShip:      mappingData.StatusShip,
		TransactionDate: order.OrderData.TransactionDate,
		Status:          order.OnlineOrderData.Status,
		NameReceiver:    order.OnlineOrderData.CustomerName,
		PhoneReceiver:   order.OnlineOrderData.CustomerPhone,
		EmailReceiver:   order.OnlineOrderData.CustomerEmail,
		Address: &schema.Address{
			Street:   order.OnlineOrderData.Street,
			Ward:     order.OnlineOrderData.Ward,
			District: order.OnlineOrderData.District,
			Province: order.OnlineOrderData.Province,
		},
		ExpectDate: order.OnlineOrderData.ExpectedDelivery,
	}, nil
}

func (c *orderServiceController) mapOrderDetailData(order *repository.OnlineOrdersResponse) *OrderDetailMappingData {
	var (
		listGoods     = make([]*schema.OrderGoodsResponse, 0, len(order.OrderGoods))
		goodsNum      = 0
		totalDiscount = 0 // total discount of an order
		statusShip    = make([]*schema.GetListOrderStateResponse, 0, len(order.ShippingState))
	)

	// append data to listGoods and assign value for total data
	for _, goods := range order.OrderGoods {
		listGoods = append(listGoods, &schema.OrderGoodsResponse{
			GoodsId:   goods.GoodsCode,
			Image:     goods.Image,
			Name:      goods.GoodsName,
			UnitPrice: goods.UnitPrice,
			Price:     goods.TotalPrice,
			Quantity:  goods.Quantity,
			Size:      goods.GoodsSize,
			Color:     goods.GoodsColor,
			Discount:  goods.Promotion,
		})
		goodsNum += goods.Quantity
		totalDiscount += int(float32(goods.UnitPrice)*goods.Promotion) * goods.Quantity
	}

	// map status ship
	for _, state := range order.ShippingState {
		statusShip = append(statusShip, &schema.GetListOrderStateResponse{
			State: state.State,
			Time:  state.StateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &OrderDetailMappingData{
		ListGoods:     listGoods,
		GoodsNum:      goodsNum,
		TotalDiscount: totalDiscount,
		StatusShip:    statusShip,
	}
}
