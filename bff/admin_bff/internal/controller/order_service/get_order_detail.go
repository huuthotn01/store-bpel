package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/order_service"
)

func (c *orderBffController) GetOrderDetail(ctx context.Context, request *order_service.GetOrderDetailRequest) (*order_service.GetOrderDetailResponseData, error) {
	order, err := c.orderAdapter.GetOrderDetail(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	var (
		listGoods        = make([]*order_service.OrderGoodsResponse, 0, len(order.ListGoods))
		onlineOrderData  *order_service.OnlineOrderData
		offlineOrderData *order_service.OfflineOrderData
	)
	for _, goods := range order.ListGoods {
		listGoods = append(listGoods, &order_service.OrderGoodsResponse{
			GoodsId:   goods.GoodsId,
			Image:     goods.Image,
			Name:      goods.Name,
			UnitPrice: goods.UnitPrice,
			Price:     goods.Price,
			Tax:       goods.Tax,
			Quantity:  goods.Quantity,
			Size:      goods.Size,
			Color:     goods.Color,
			Discount:  goods.Discount,
		})
	}

	if order.IsOnline {
		onlineOrderData = &order_service.OnlineOrderData{
			PaymentMethod: order.OnlineOrderData.PaymentMethod,
			CustomerId:    order.OnlineOrderData.CustomerId,
			IsCompleted:   order.OnlineOrderData.IsCompleted,
			ShipFee:       order.OnlineOrderData.ShipFee,
			ExpectDate:    order.OnlineOrderData.ExpectDate,
			Status:        order.OnlineOrderData.Status,
			NameReceiver:  order.OnlineOrderData.NameReceiver,
			PhoneReceiver: order.OnlineOrderData.PhoneReceiver,
			EmailReceiver: order.OnlineOrderData.EmailReceiver,
			Address: &order_service.Address{
				Street:   order.OnlineOrderData.Address.Street,
				Ward:     order.OnlineOrderData.Address.Ward,
				District: order.OnlineOrderData.Address.District,
				Province: order.OnlineOrderData.Address.Province,
			},
		}
	} else {
		offlineOrderData = &order_service.OfflineOrderData{
			StaffId:  order.OfflineOrderData.StaffId,
			BranchId: order.OfflineOrderData.BranchId,
		}
	}

	return &order_service.GetOrderDetailResponseData{
		OrderId:          order.OrderId,
		OrderCode:        order.OrderCode,
		ListGoods:        listGoods,
		TotalPrice:       order.TotalPrice,
		TotalGoods:       order.TotalGoods,
		TotalDiscount:    order.TotalDiscount,
		TotalOrder:       order.TotalOrder,
		TransactionDate:  order.TransactionDate,
		IsOnline:         order.IsOnline,
		OnlineOrderData:  onlineOrderData,
		OfflineOrderData: offlineOrderData,
	}, nil
}
