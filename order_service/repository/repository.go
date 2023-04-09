package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type IOrderServiceRepository interface {
	GetOnlineOrders(ctx context.Context) ([]*OnlineOrderJoiningResponse, error)
	GetOfflineOrders(ctx context.Context) ([]*OfflineOrderJoiningResponse, error)
	GetOnlineOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error)
	GetOrderByOrderId(ctx context.Context, orderId int) (*OrdersModel, error)
	GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*GoodsModel, error)
	GetOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersModel, error)
	GetOnlineOrderDetail(ctx context.Context, privateOrderId int) (*OnlineOrdersResponse, error)
	GetOfflineOrderDetail(ctx context.Context, privateOrderId int) (*OfflineOrdersResponse, error)
	GetPrivateOrderCode(ctx context.Context, orderId string) (int, error)
	GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error)
	GetOnlineOrderByOrderId(ctx context.Context, orderId int) (*OnlineOrdersModel, error)
	GetOfflineOrderByOrderId(ctx context.Context, orderId int) (*StoreOrdersModel, error)

	CreateOnlineOrder(ctx context.Context, data *OnlineOrdersData) error
	CreateOfflineOrder(ctx context.Context, data *OfflineOrdersData) error
	UpdateOrderState(ctx context.Context, orderState *OnlineOrderStateData) error
}

func NewRepository(db *gorm.DB) IOrderServiceRepository {
	return &orderServiceRepository{
		db:                    db,
		goodsTableName:        "goods",
		ordersTableName:       "orders",
		onlineOrdersTableName: "online_orders",
		storeOrdersTableName:  "store_orders",
		orderStateTableName:   "order_state",
	}
}

var ErrOrderNotFound = errors.New("given order not found")

type (
	OnlineOrdersData struct {
		PublicOrderCode string
		TransactionDate string
		TotalPrice      int
		OnlineOrder     *OnlineOrdersModel
		Goods           []*GoodsModel
	}

	OfflineOrdersData struct {
		PublicOrderCode string
		TransactionDate string
		TotalPrice      int
		StaffId         string
		BranchId        string
		Goods           []*GoodsModel
	}

	OnlineOrdersResponse struct {
		OrderGoods      []*GoodsModel
		OrderData       *OrdersModel
		OnlineOrderData *OnlineOrdersModel
		ShippingState   []*OrderStateModel
	}

	OfflineOrdersResponse struct {
		OrderGoods       []*GoodsModel
		OrderData        *OrdersModel
		OfflineOrderData *StoreOrdersModel
	}

	OnlineOrderStateData struct {
		OrderState   *OrderStateModel
		StatusNumber int
	}

	OnlineOrderJoiningResponse struct {
		OrderCode        int
		PublicOrderCode  string
		GoodsCode        string
		Image            string
		GoodsName        string
		UnitPrice        int
		Price            int
		Tax              float32
		Quantity         int
		GoodsSize        string
		GoodsColor       string
		Promotion        float32
		TotalPrice       int
		TransactionDate  string
		PaymentMethod    string
		CustomerId       string
		ShippingFee      int
		ExpectedDelivery string
		Status           int
		CustomerName     string
		CustomerPhone    string
		CustomerEmail    string
		Street           string
		Ward             string
		District         string
		Province         string
	}

	OfflineOrderJoiningResponse struct {
		OrderCode       int
		PublicOrderCode string
		GoodsCode       string
		Image           string
		GoodsName       string
		UnitPrice       int
		Price           int
		Tax             float32
		Quantity        int
		GoodsSize       string
		GoodsColor      string
		Promotion       float32
		TotalPrice      int
		TransactionDate string
		StaffId         string
		StoreCode       string
	}
)

/*
	***************
	| GET METHODS |
	***************
*/
func (r *orderServiceRepository) GetOnlineOrders(ctx context.Context) ([]*OnlineOrderJoiningResponse, error) {
	var result []*OnlineOrderJoiningResponse
	query := r.db.WithContext(ctx).Table(r.onlineOrdersTableName).
		Joins("goods on online_orders.order_code = goods.order_code").
		Joins("orders on online_orders.order_code = orders.order_code").
		Select("orders.*, online_orders.*, goods.goods_code, goods.image, goods.goods_name, goods.unit_price, goods.total_price as price, goods.tax, goods.quantity, goods.goods_size, goods.goods_color, goods.promotion")
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetOfflineOrders(ctx context.Context) ([]*OfflineOrderJoiningResponse, error) {
	var result []*OfflineOrderJoiningResponse
	query := r.db.WithContext(ctx).Table(r.storeOrdersTableName).
		Joins("goods on store_orders.order_code = goods.order_code").
		Joins("orders on store_orders.order_code = orders.order_code").
		Select("orders.*, goods.goods_code, goods.image, goods.goods_name, goods.unit_price, goods.total_price as price, goods.tax, goods.quantity, goods.goods_size, goods.goods_color, goods.promotion, store_orders.*")
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetOnlineOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersResponse, error) {
	// get orders made by customer
	orders, err := r.GetOrdersByCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}

	resp := make([]*OnlineOrdersResponse, 0, len(orders))

	for _, order := range orders {
		detail, err := r.GetOnlineOrderDetail(ctx, order.OrderCode)
		if err != nil {
			return nil, err
		}

		resp = append(resp, detail)
	}

	return resp, nil
}

func (r *orderServiceRepository) GetOnlineOrderDetail(ctx context.Context, privateOrderId int) (*OnlineOrdersResponse, error) {
	// get order data
	order, err := r.GetOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	// get online order data
	onlineOrder, err := r.GetOnlineOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get order goods
	goods, err := r.GetOrderGoodsByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get order state
	state, err := r.GetOrderState(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	return &OnlineOrdersResponse{
		OrderData:       order,
		OrderGoods:      goods,
		ShippingState:   state,
		OnlineOrderData: onlineOrder,
	}, nil
}

func (r *orderServiceRepository) GetOfflineOrderDetail(ctx context.Context, privateOrderId int) (*OfflineOrdersResponse, error) {
	// get order data
	order, err := r.GetOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get offline order data
	offlineOrder, err := r.GetOfflineOrderByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	// get order goods
	goods, err := r.GetOrderGoodsByOrderId(ctx, privateOrderId)
	if err != nil {
		return nil, err
	}

	return &OfflineOrdersResponse{
		OrderData:        order,
		OrderGoods:       goods,
		OfflineOrderData: offlineOrder,
	}, nil
}

func (r *orderServiceRepository) GetOrdersByCustomer(ctx context.Context, customerId string) ([]*OnlineOrdersModel, error) {
	var result []*OnlineOrdersModel
	query := r.db.WithContext(ctx).Table(r.onlineOrdersTableName).Where("customer_id = ?", customerId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetPrivateOrderCode(ctx context.Context, orderId string) (int, error) {
	var result int
	query := r.db.WithContext(ctx).Table(r.ordersTableName).Where("public_order_code = ?", orderId).Select("order_code")
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

func (r *orderServiceRepository) GetOrderByOrderId(ctx context.Context, orderId int) (*OrdersModel, error) {
	var result *OrdersModel
	query := r.db.WithContext(ctx).Table(r.ordersTableName).Where("order_code = ?", orderId)
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOnlineOrderByOrderId(ctx context.Context, orderId int) (*OnlineOrdersModel, error) {
	var result *OnlineOrdersModel
	query := r.db.WithContext(ctx).Table(r.onlineOrdersTableName).Where("order_code = ?", orderId)
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOfflineOrderByOrderId(ctx context.Context, orderId int) (*StoreOrdersModel, error) {
	var result *StoreOrdersModel
	query := r.db.WithContext(ctx).Table(r.storeOrdersTableName).Where("order_code = ?", orderId)
	return result, query.First(&result).Error
}

func (r *orderServiceRepository) GetOrderState(ctx context.Context, orderId int) ([]*OrderStateModel, error) {
	var result []*OrderStateModel
	query := r.db.WithContext(ctx).Table(r.orderStateTableName).Where("order_code = ?", orderId)
	return result, query.Find(&result).Error
}

/*
	***************
	| UPDATE METHODS |
	***************
*/
func (r *orderServiceRepository) CreateOnlineOrder(ctx context.Context, data *OnlineOrdersData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// add data to orders table
		orderModel := &OrdersModel{
			TransactionDate: data.TransactionDate,
			TotalPrice:      data.TotalPrice,
			PublicOrderCode: data.PublicOrderCode,
		}
		err := tx.Table(r.ordersTableName).Create(orderModel).Error
		if err != nil {
			return err
		}

		// add data to goods table
		for _, g := range data.Goods {
			g.OrderCode = orderModel.OrderCode
		}
		err = tx.Table(r.goodsTableName).Create(data.Goods).Error
		if err != nil {
			return err
		}

		// add data to online_orders table
		data.OnlineOrder.OrderCode = orderModel.OrderCode
		err = tx.Table(r.onlineOrdersTableName).Create(data.OnlineOrder).Error
		return err
	})
}

func (r *orderServiceRepository) CreateOfflineOrder(ctx context.Context, data *OfflineOrdersData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// add data to orders table
		orderModel := &OrdersModel{
			TransactionDate: data.TransactionDate,
			TotalPrice:      data.TotalPrice,
			PublicOrderCode: data.PublicOrderCode,
		}
		err := tx.Table(r.ordersTableName).Create(orderModel).Error
		if err != nil {
			return err
		}

		// add data to goods table
		for _, g := range data.Goods {
			g.OrderCode = orderModel.OrderCode
		}
		err = tx.Table(r.goodsTableName).Create(data.Goods).Error
		if err != nil {
			return err
		}

		// add data to store orders table
		storeOrdersModel := &StoreOrdersModel{
			OrderCode: orderModel.OrderCode,
			StoreCode: data.BranchId,
			StaffId:   data.StaffId,
		}
		return tx.Table(r.storeOrdersTableName).Create(&storeOrdersModel).Error
	})
}

func (r *orderServiceRepository) UpdateOrderState(ctx context.Context, orderState *OnlineOrderStateData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// add new state of order
		err := tx.Table(r.orderStateTableName).Select("order_code", "state").Create(&orderState.OrderState).Error
		if err != nil {
			return err
		}

		// update order status number in online order table
		return tx.Table(r.onlineOrdersTableName).Where("order_code = ?", orderState.OrderState.OrderCode).
			Update("status", orderState.StatusNumber).Error
	})
}
