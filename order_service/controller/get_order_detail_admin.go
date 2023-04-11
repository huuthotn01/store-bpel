package controller

import (
	"context"
	"errors"
	"store-bpel/order_service/repository"
	"store-bpel/order_service/schema"
)

type OrderGoodsAndMoneyData struct {
	ListGoods     []*schema.OrderGoodsResponse
	TotalGoods    int
	TotalDiscount int
}

func (c *orderServiceController) GetOrderDetailAdmin(ctx context.Context, orderId int) (*schema.GetOrderDetailAdminResponseData, error) {
	onlineOrder, err := c.repository.GetOnlineOrderDetail(ctx, orderId)
	if err == nil {
		// case online order
		orderGoodsMoneyData := c.mapListGoods(onlineOrder.OrderGoods)
		return &schema.GetOrderDetailAdminResponseData{
			OrderId:       onlineOrder.OrderData.OrderCode,
			OrderCode:     onlineOrder.OrderData.PublicOrderCode,
			ListGoods:     orderGoodsMoneyData.ListGoods,
			TotalPrice:    onlineOrder.OrderData.TotalPrice,
			TotalGoods:    orderGoodsMoneyData.TotalGoods,
			TotalDiscount: orderGoodsMoneyData.TotalDiscount,
			// total order = total price of orders' items + shipping fee - discount
			TotalOrder:      onlineOrder.OrderData.TotalPrice + onlineOrder.OnlineOrderData.ShippingFee,
			TransactionDate: onlineOrder.OrderData.TransactionDate,
			IsOnline:        true,
			OnlineOrderData: &schema.OnlineOrderData{
				PaymentMethod: onlineOrder.OnlineOrderData.PaymentMethod,
				CustomerId:    onlineOrder.OnlineOrderData.CustomerId,
				IsCompleted:   onlineOrder.OnlineOrderData.Status == 4,
				ShipFee:       onlineOrder.OnlineOrderData.ShippingFee,
				ExpectDate:    onlineOrder.OnlineOrderData.ExpectedDelivery,
				Status:        onlineOrder.OnlineOrderData.Status,
				NameReceiver:  onlineOrder.OnlineOrderData.CustomerName,
				PhoneReceiver: onlineOrder.OnlineOrderData.CustomerPhone,
				EmailReceiver: onlineOrder.OnlineOrderData.CustomerEmail,
				Address: &schema.Address{
					Street:   onlineOrder.OnlineOrderData.Street,
					Ward:     onlineOrder.OnlineOrderData.Ward,
					District: onlineOrder.OnlineOrderData.District,
					Province: onlineOrder.OnlineOrderData.Province,
				},
			},
		}, err
	}

	// order not exists
	if errors.Is(err, repository.ErrOrderNotFound) {
		return nil, err
	}

	offlineOrder, err := c.repository.GetOfflineOrderDetail(ctx, orderId)
	if err != nil {
		return nil, err
	}

	// in case offline order
	orderGoodsMoneyData := c.mapListGoods(offlineOrder.OrderGoods)
	return &schema.GetOrderDetailAdminResponseData{
		OrderId:         offlineOrder.OrderData.OrderCode,
		OrderCode:       offlineOrder.OrderData.PublicOrderCode,
		ListGoods:       orderGoodsMoneyData.ListGoods,
		TotalPrice:      offlineOrder.OrderData.TotalPrice,
		TotalGoods:      orderGoodsMoneyData.TotalGoods,
		TotalDiscount:   orderGoodsMoneyData.TotalDiscount,
		TotalOrder:      offlineOrder.OrderData.TotalPrice,
		TransactionDate: offlineOrder.OrderData.TransactionDate,
		IsOnline:        false,
		OfflineOrderData: &schema.OfflineOrderData{
			StaffId:  offlineOrder.OfflineOrderData.StaffId,
			BranchId: offlineOrder.OfflineOrderData.StoreCode,
		},
	}, nil
}

func (c *orderServiceController) mapListGoods(data []*repository.GoodsModel) *OrderGoodsAndMoneyData {
	var (
		listGoods     = make([]*schema.OrderGoodsResponse, 0, len(data))
		goodsNum      = 0
		totalDiscount = 0 // total discount of an order
	)
	for _, goods := range data {
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
		totalDiscount += int(float32(goods.TotalPrice) * goods.Promotion)
	}

	return &OrderGoodsAndMoneyData{
		ListGoods:     listGoods,
		TotalGoods:    goodsNum,
		TotalDiscount: totalDiscount,
	}
}
