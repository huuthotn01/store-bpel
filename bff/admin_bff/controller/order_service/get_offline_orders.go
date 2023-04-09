package order_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/order_service"
)

func (c *orderBffController) GetOfflineOrders(ctx context.Context) ([]*order_service.GetOfflineOrdersResponseData, error) {
	orders, err := c.orderAdapter.GetOfflineOrders(ctx)
	if err != nil {
		return nil, err
	}

	respOrders := make([]*order_service.GetOfflineOrdersResponseData, 0, len(orders))
	for _, order := range orders {
		listGoods := make([]*order_service.OrderGoodsResponse, 0, len(order.ListGoods))
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

		respOrders = append(respOrders, &order_service.GetOfflineOrdersResponseData{
			OrderId:         order.OrderId,
			OrderCode:       order.OrderCode,
			ListGoods:       listGoods,
			TotalPrice:      order.TotalPrice,
			TotalGoods:      order.TotalGoods,
			TotalDiscount:   order.TotalDiscount,
			TotalOrder:      order.TotalOrder,
			TransactionDate: order.TransactionDate,
			OfflineOrderData: &order_service.OfflineOrderData{
				StaffId:  order.OfflineOrderData.StaffId,
				BranchId: order.OfflineOrderData.BranchId,
			},
		})
	}

	return respOrders, nil
}
