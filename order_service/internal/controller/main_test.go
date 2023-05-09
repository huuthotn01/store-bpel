package controller

import (
	"context"
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
	return nil, nil
}

func (t *testRepo) GetOfflineOrders(ctx context.Context) ([]*repository.OfflineOrderJoiningResponse, error) {
	return nil, nil
}

func (t *testRepo) GetOnlineOrdersByCustomer(ctx context.Context, customerId string) ([]*repository.OnlineOrdersResponse, error) {
	return nil, nil
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
	return nil, nil
}

func (t *testRepo) GetOfflineOrderDetail(ctx context.Context, privateOrderId int) (*repository.OfflineOrdersResponse, error) {
	return nil, nil
}

func (t *testRepo) GetPrivateOrderCode(ctx context.Context, orderId string) (int, error) {
	return 5, nil
}

func (t *testRepo) GetOrderState(ctx context.Context, orderId int) ([]*repository.OrderStateModel, error) {
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
