package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"os"
	"store-bpel/goods_service/schema"
	"store-bpel/order_service/internal/repository"
	"testing"
	"time"
)

var (
	testRepository TestRepository
	testGoods      TestGoodsAdapter
	testKafka      TestKafkaAdapter
)

// MOCK KAFKA ADAPTER
type TestKafkaAdapter interface {
	Publish(ctx context.Context, topic string, msg []byte) error
}

type testKafkaAdapter struct {
}

func NewTestKafkaAdapter() TestKafkaAdapter {
	return &testKafkaAdapter{}
}

func (a *testKafkaAdapter) Publish(ctx context.Context, topic string, msg []byte) error {
	return nil
}

// MOCK GOODS ADAPTER
type TestGoodsAdapter interface {
	GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error)
}

type testGoodsAdapter struct {
}

func NewTestStaffAdapter() TestGoodsAdapter {
	return &testGoodsAdapter{}
}

func (t *testGoodsAdapter) GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error) {
	if productId == "invalid-goods-id" {
		return nil, errors.New("some random error")
	}
	return &schema.GetGoodsDefaultResponseData{
		Name:        "Goods Name",
		GoodsType:   "Type 1",
		GoodsGender: 3,
		UnitCost:    1000,
	}, nil
}

// MOCK REPOSITORY
type TestRepository interface {
	GetOnlineOrders(ctx context.Context) ([]*repository.OnlineOrderJoiningResponse, error)
	GetOfflineOrders(ctx context.Context) ([]*repository.OfflineOrderJoiningResponse, error)
	GetOnlineOrdersByCustomer(ctx context.Context, customerId string) ([]*repository.OnlineOrdersResponse, error)
	GetOrderByOrderId(ctx context.Context, orderId int) (*repository.OrdersModel, error)
	GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*repository.GoodsModel, error)
	GetOrdersByCustomer(ctx context.Context, customerId string) ([]*repository.OnlineOrdersModel, error)
	GetOnlineOrderDetail(ctx context.Context, privateOrderId int) (*repository.OnlineOrdersResponse, error)
	GetOfflineOrderDetail(ctx context.Context, privateOrderId int) (*repository.OfflineOrdersResponse, error)
	GetPrivateOrderCode(ctx context.Context, orderId string) (int, error)
	GetOrderState(ctx context.Context, orderId int) ([]*repository.OrderStateModel, error)
	GetOnlineOrderByOrderId(ctx context.Context, orderId int) (*repository.OnlineOrdersModel, error)
	GetOfflineOrderByOrderId(ctx context.Context, orderId int) (*repository.StoreOrdersModel, error)
	GetBestGoods(ctx context.Context) ([]string, error)

	CreateOnlineOrder(ctx context.Context, data *repository.OnlineOrdersData) error
	CreateOfflineOrder(ctx context.Context, data *repository.OfflineOrdersData) error
	UpdateOrderState(ctx context.Context, orderState *repository.OnlineOrderStateData) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetOnlineOrders(ctx context.Context) ([]*repository.OnlineOrderJoiningResponse, error) {
	return []*repository.OnlineOrderJoiningResponse{
		{
			OrderCode:        5,
			PublicOrderCode:  "MNOPQR",
			GoodsCode:        "goods-1",
			GoodsName:        "Goods One",
			UnitPrice:        10000,
			Price:            20000,
			Quantity:         2,
			Tax:              0.1,
			ShippingFee:      5000,
			GoodsColor:       "red",
			GoodsSize:        "XL",
			Promotion:        0.2,
			TotalPrice:       40000,
			TransactionDate:  "2023-01-01",
			PaymentMethod:    "COD",
			CustomerId:       "customer-1",
			ExpectedDelivery: "2023-01-06",
			Status:           3,
			CustomerName:     "HTTN",
			CustomerPhone:    "0123456789",
			CustomerEmail:    "httn@gmail.com",
			Street:           "THT",
			Ward:             "Ward 11",
			District:         "District 10",
			Province:         "Ho Chi Minh City",
		},
		{
			OrderCode:        6,
			PublicOrderCode:  "AAAAAA",
			GoodsCode:        "goods-2",
			GoodsName:        "Goods Two",
			UnitPrice:        20000,
			Price:            22000,
			Quantity:         1,
			Tax:              0.1,
			ShippingFee:      3000,
			GoodsColor:       "blue",
			GoodsSize:        "XXL",
			Promotion:        0.2,
			TotalPrice:       22000,
			TransactionDate:  "2023-01-02",
			PaymentMethod:    "momo",
			CustomerId:       "customer-2",
			ExpectedDelivery: "2023-01-08",
			Status:           4,
			CustomerName:     "Huu Tho",
			CustomerPhone:    "0111111111",
			CustomerEmail:    "tho@gmail.com",
			Street:           "LTK",
			Ward:             "P.11",
			District:         "Q.10",
			Province:         "HCMC",
		},
	}, nil
}

func (t *testRepo) GetOfflineOrders(ctx context.Context) ([]*repository.OfflineOrderJoiningResponse, error) {
	return []*repository.OfflineOrderJoiningResponse{
		{
			OrderCode:       3,
			PublicOrderCode: "MNOPQR",
			GoodsCode:       "goods-11",
			GoodsName:       "Goods Eleven",
			UnitPrice:       10000,
			Price:           20000,
			Quantity:        2,
			Tax:             0.1,
			GoodsColor:      "red",
			GoodsSize:       "XL",
			Promotion:       0.2,
			TotalPrice:      40000,
			TransactionDate: "2023-01-01",
			StaffId:         "staff-1",
			StoreCode:       "store-1",
		},
		{
			OrderCode:       4,
			PublicOrderCode: "AAAAAA",
			GoodsCode:       "goods-12",
			GoodsName:       "Goods Twelve",
			UnitPrice:       20000,
			Price:           22000,
			Quantity:        1,
			Tax:             0.1,
			GoodsColor:      "blue",
			GoodsSize:       "XXL",
			Promotion:       0.2,
			TotalPrice:      22000,
			TransactionDate: "2023-01-02",
			StaffId:         "staff-3",
			StoreCode:       "store-2",
		},
	}, nil
}

func (t *testRepo) GetOnlineOrdersByCustomer(ctx context.Context, customerId string) ([]*repository.OnlineOrdersResponse, error) {
	return []*repository.OnlineOrdersResponse{
		{
			OrderGoods: []*repository.GoodsModel{
				{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					GoodsName:  "Goods One",
					OrderCode:  5,
					Quantity:   3,
					UnitPrice:  5000,
					TotalPrice: 5500,
					Tax:        0.1,
					Promotion:  0.1,
				},
				{
					GoodsCode:  "goods-2",
					GoodsColor: "yellow",
					GoodsSize:  "S",
					GoodsName:  "Goods Two",
					OrderCode:  5,
					Quantity:   1,
					UnitPrice:  5000,
					TotalPrice: 5500,
					Tax:        0.1,
					Promotion:  0.1,
				},
			},
			OrderData: &repository.OrdersModel{
				OrderCode:       5,
				TransactionDate: "2023-01-01",
				TotalPrice:      11000,
				PublicOrderCode: "ABCDEF",
			},
			OnlineOrderData: &repository.OnlineOrdersModel{
				OrderCode:        5,
				ExpectedDelivery: "2023-01-06",
				ShippingFee:      1000,
				CustomerId:       "customer-1",
				PaymentMethod:    "COD",
				Street:           "LTK",
				Ward:             "11",
				District:         "10",
				Province:         "HCMC",
				CustomerName:     "HTTN",
				CustomerPhone:    "0123456789",
				CustomerEmail:    "httn@gmail.com",
				Status:           4,
			},
			ShippingState: []*repository.OrderStateModel{
				{
					OrderCode: 5,
					State:     "Packed by seller",
					StateTime: time.Date(2023, 1, 1, 15, 16, 17, 0, time.Local),
				},
				{
					OrderCode: 5,
					State:     "Picked by shipper",
					StateTime: time.Date(2023, 1, 2, 8, 9, 10, 0, time.Local),
				},
			},
		},
	}, nil
}

func (t *testRepo) GetOrderByOrderId(ctx context.Context, orderId int) (*repository.OrdersModel, error) {
	return nil, nil
}

func (t *testRepo) GetOrderGoodsByOrderId(ctx context.Context, orderId int) ([]*repository.GoodsModel, error) {
	return nil, nil
}

func (t *testRepo) GetOrdersByCustomer(ctx context.Context, customerId string) ([]*repository.OnlineOrdersModel, error) {
	return nil, nil
}

func (t *testRepo) GetOnlineOrderDetail(ctx context.Context, privateOrderId int) (*repository.OnlineOrdersResponse, error) {
	if privateOrderId == 10 {
		return nil, errors.New("some random error")
	}
	if privateOrderId == 3 || privateOrderId == 4 || privateOrderId == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &repository.OnlineOrdersResponse{
		OrderGoods: []*repository.GoodsModel{
			{
				GoodsCode:  "goods-1",
				GoodsColor: "red",
				GoodsSize:  "XL",
				GoodsName:  "Goods One",
				OrderCode:  5,
				Quantity:   3,
				UnitPrice:  5000,
				TotalPrice: 5500,
				Tax:        0.1,
				Promotion:  0.1,
			},
			{
				GoodsCode:  "goods-2",
				GoodsColor: "yellow",
				GoodsSize:  "S",
				GoodsName:  "Goods Two",
				OrderCode:  5,
				Quantity:   1,
				UnitPrice:  5000,
				TotalPrice: 5500,
				Tax:        0.1,
				Promotion:  0.1,
			},
		},
		OrderData: &repository.OrdersModel{
			OrderCode:       5,
			TransactionDate: "2023-01-01",
			TotalPrice:      11000,
			PublicOrderCode: "ABCDEF",
		},
		OnlineOrderData: &repository.OnlineOrdersModel{
			OrderCode:        5,
			ExpectedDelivery: "2023-01-06",
			ShippingFee:      1000,
			CustomerId:       "customer-1",
			PaymentMethod:    "COD",
			Street:           "LTK",
			Ward:             "11",
			District:         "10",
			Province:         "HCMC",
			CustomerName:     "HTTN",
			CustomerPhone:    "0123456789",
			CustomerEmail:    "httn@gmail.com",
			Status:           4,
		},
		ShippingState: []*repository.OrderStateModel{
			{
				OrderCode: 5,
				State:     "Packed by seller",
				StateTime: time.Date(2023, 1, 1, 15, 16, 17, 0, time.Local),
			},
			{
				OrderCode: 5,
				State:     "Picked by shipper",
				StateTime: time.Date(2023, 1, 2, 8, 9, 10, 0, time.Local),
			},
		},
	}, nil
}

func (t *testRepo) GetOfflineOrderDetail(ctx context.Context, privateOrderId int) (*repository.OfflineOrdersResponse, error) {
	if privateOrderId == 0 {
		return nil, errors.New("some random error")
	}
	if privateOrderId == 5 || privateOrderId == 6 {
		return nil, gorm.ErrRecordNotFound
	}
	return &repository.OfflineOrdersResponse{
		OrderGoods: []*repository.GoodsModel{
			{
				GoodsCode:  "goods-1",
				GoodsColor: "red",
				GoodsSize:  "XL",
				GoodsName:  "Goods One",
				OrderCode:  3,
				Quantity:   3,
				UnitPrice:  5000,
				TotalPrice: 5500,
				Tax:        0.1,
				Promotion:  0.1,
			},
			{
				GoodsCode:  "goods-2",
				GoodsColor: "yellow",
				GoodsSize:  "S",
				GoodsName:  "Goods Two",
				OrderCode:  3,
				Quantity:   1,
				UnitPrice:  5000,
				TotalPrice: 5500,
				Tax:        0.1,
				Promotion:  0.1,
			},
		},
		OrderData: &repository.OrdersModel{
			OrderCode:       3,
			TransactionDate: "2023-01-01",
			TotalPrice:      11000,
			PublicOrderCode: "ABCDEF",
		},
		OfflineOrderData: &repository.StoreOrdersModel{
			OrderCode: 3,
			StoreCode: "store-1",
			StaffId:   "staff-1",
		},
	}, nil
}

func (t *testRepo) GetPrivateOrderCode(ctx context.Context, orderId string) (int, error) {
	if orderId == "invalid-order" {
		return 0, errors.New("some random error")
	}
	if orderId == "invalid-order-state" {
		return 9, nil
	}
	return 5, nil
}

func (t *testRepo) GetOrderState(ctx context.Context, orderId int) ([]*repository.OrderStateModel, error) {
	if orderId == 9 {
		return nil, errors.New("some random error")
	}
	return []*repository.OrderStateModel{
		{
			OrderCode: 5,
			State:     "Packed by seller",
			StateTime: time.Date(2023, 1, 1, 7, 0, 0, 0, time.Local),
		},
		{
			OrderCode: 5,
			State:     "Picked from warehouse",
			StateTime: time.Date(2023, 1, 3, 13, 0, 0, 0, time.Local),
		},
		{
			OrderCode: 5,
			State:     "Delivered",
			StateTime: time.Date(2023, 1, 6, 9, 0, 0, 0, time.Local),
		},
	}, nil
}

func (t *testRepo) GetOnlineOrderByOrderId(ctx context.Context, orderId int) (*repository.OnlineOrdersModel, error) {
	return nil, nil
}

func (t *testRepo) GetOfflineOrderByOrderId(ctx context.Context, orderId int) (*repository.StoreOrdersModel, error) {
	return nil, nil
}

func (t *testRepo) GetBestGoods(ctx context.Context) ([]string, error) {
	return []string{"goods-1", "goods-2"}, nil
}

func (t *testRepo) CreateOnlineOrder(ctx context.Context, data *repository.OnlineOrdersData) error {
	return nil
}

func (t *testRepo) CreateOfflineOrder(ctx context.Context, data *repository.OfflineOrdersData) error {
	if data.StaffId == "invalid-offline-order" {
		return errors.New("some random error")
	}
	return nil
}

func (t *testRepo) UpdateOrderState(ctx context.Context, orderState *repository.OnlineOrderStateData) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testGoods = NewTestStaffAdapter()
	testKafka = NewTestKafkaAdapter()

	exitVal := m.Run()

	os.Exit(exitVal)
}
