package controller

import (
	"context"
	"os"
	event_schema "store-bpel/event_service/schema"
	"store-bpel/goods_service/internal/repository"
	"store-bpel/warehouse_service/schema"
	"testing"
)

var (
	testRepository TestRepository
	testOrder      TestOrderAdapter
	testEvent      TestEventAdapter
	testWarehouse  TestWarehouseAdapter
)

// MOCK EVENT ADAPTER
type TestEventAdapter interface {
	GetEventByGoods(ctx context.Context, goodsId string) ([]*event_schema.GetEventByGoodsData, error)
}

type testEventAdapter struct {
}

func NewTestEventAdapter() TestEventAdapter {
	return &testEventAdapter{}
}

func (a *testEventAdapter) GetEventByGoods(ctx context.Context, goodsId string) ([]*event_schema.GetEventByGoodsData, error) {
	return nil, nil
}

// MOCK ORDER ADAPTER
type TestOrderAdapter interface {
	GetBestSellingGoods(ctx context.Context) ([]string, error)
}

type testOrderAdapter struct {
}

func NewTestOrderAdapter() TestOrderAdapter {
	return &testOrderAdapter{}
}

func (t *testOrderAdapter) GetBestSellingGoods(ctx context.Context) ([]string, error) {
	return []string{"goods-1"}, nil
}

// MOCK WAREHOUSE ADAPTER
type TestWarehouseAdapter interface {
	GetWarehouse(ctx context.Context) (*schema.UpdateResponse, error)
}

type testWarehouseAdapter struct {
}

func NewTestWarehouseAdapter() TestWarehouseAdapter {
	return &testWarehouseAdapter{}
}

func (t *testWarehouseAdapter) GetWarehouse(ctx context.Context) (*schema.UpdateResponse, error) {
	return nil, nil
}

// MOCK REPOSITORY
type TestRepository interface {
	FilterGoods(ctx context.Context, name string, newAdded bool) ([]string, error)
	GetGoods(ctx context.Context) ([]*repository.GoodsModel, error)
	GetGoodsDefault(ctx context.Context, limit, offset int) ([]string, error)
	GetImages(ctx context.Context) ([]*repository.GoodsImg, error)
	GetGoodsImages(ctx context.Context, goodsId string) ([]*repository.GoodsImg, error)
	GetGoodsImageUrls(ctx context.Context, goodsId string) ([]string, error)
	GetDetailGoods(ctx context.Context, goodsId string) ([]*repository.GoodsModel, error)
	AddGoods(ctx context.Context, data []*repository.GoodsModel) error
	UpdateGoods(ctx context.Context, data []*repository.GoodsModel) error
	UpdateGoodsIsForSaleToNo(ctx context.Context, goodsId string) error
	GetGoodsInWHData(ctx context.Context, data *repository.GoodsInWh) ([]*repository.GoodsInWh, error)
	UpdateGoodsInWHInOut(ctx context.Context, data *repository.GoodsInWh) error
	UpdateGoodsInWHTransfer(ctx context.Context, data *repository.GoodsInWh, fromWH, toWH string) error
	GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*repository.GoodsInWh, error)
	AddGoodsImage(ctx context.Context, data *repository.GoodsImg) error
	DeleteGoodsImage(ctx context.Context, url string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) FilterGoods(ctx context.Context, name string, newAdded bool) ([]string, error) {
	return []string{"goods-1"}, nil
}

func (t *testRepo) GetGoods(ctx context.Context) ([]*repository.GoodsModel, error) {
	return []*repository.GoodsModel{
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsSize:  "XL",
			GoodsName:  "Goods One",
		},
		{
			GoodsCode:  "goods-1",
			GoodsColor: "yellow",
			GoodsSize:  "XXL",
			GoodsName:  "Goods One",
		},
	}, nil
}

func (t *testRepo) GetGoodsDefault(ctx context.Context, limit, offset int) ([]string, error) {
	return []string{"goods-1"}, nil
}

func (t *testRepo) GetImages(ctx context.Context) ([]*repository.GoodsImg, error) {
	return []*repository.GoodsImg{
		{
			GoodsCode:  "goods-1",
			GoodsColor: "yellow",
			GoodsImg:   "url-1",
		},
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsImg:   "url-2",
		},
	}, nil
}

func (t *testRepo) GetGoodsImages(ctx context.Context, goodsId string) ([]*repository.GoodsImg, error) {
	return []*repository.GoodsImg{
		{
			GoodsCode:  "goods-1",
			GoodsColor: "yellow",
			GoodsImg:   "url-1",
		},
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsImg:   "url-2",
		},
	}, nil
}

func (t *testRepo) GetGoodsImageUrls(ctx context.Context, goodsId string) ([]string, error) {
	return nil, nil
}

func (t *testRepo) GetDetailGoods(ctx context.Context, goodsId string) ([]*repository.GoodsModel, error) {
	return []*repository.GoodsModel{
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsSize:  "XL",
			GoodsName:  "Goods One",
		},
		{
			GoodsCode:  "goods-1",
			GoodsColor: "yellow",
			GoodsSize:  "XXL",
			GoodsName:  "Goods One",
		},
	}, nil
}

func (t *testRepo) AddGoods(ctx context.Context, data []*repository.GoodsModel) error {
	return nil
}

func (t *testRepo) UpdateGoods(ctx context.Context, data []*repository.GoodsModel) error {
	return nil
}

func (t *testRepo) UpdateGoodsIsForSaleToNo(ctx context.Context, goodsId string) error {
	return nil
}

func (t *testRepo) GetGoodsInWHData(ctx context.Context, data *repository.GoodsInWh) ([]*repository.GoodsInWh, error) {
	return []*repository.GoodsInWh{
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsSize:  "XL",
			Quantity:   5,
			WhCode:     "warehouse-1",
		},
		{
			GoodsCode:  "goods-1",
			GoodsColor: "red",
			GoodsSize:  "XL",
			Quantity:   7,
			WhCode:     "warehouse-2",
		},
	}, nil
}

func (t *testRepo) UpdateGoodsInWHInOut(ctx context.Context, data *repository.GoodsInWh) error {
	return nil
}

func (t *testRepo) UpdateGoodsInWHTransfer(ctx context.Context, data *repository.GoodsInWh, fromWH, toWH string) error {
	return nil
}

func (t *testRepo) GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*repository.GoodsInWh, error) {
	return nil, nil
}

func (t *testRepo) AddGoodsImage(ctx context.Context, data *repository.GoodsImg) error {
	return nil
}

func (t *testRepo) DeleteGoodsImage(ctx context.Context, url string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testOrder = NewTestOrderAdapter()
	testEvent = NewTestEventAdapter()
	testWarehouse = NewTestWarehouseAdapter()

	exitVal := m.Run()

	os.Exit(exitVal)
}
