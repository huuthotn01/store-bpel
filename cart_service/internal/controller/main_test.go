package controller

import (
	"context"
	"errors"
	"os"
	"store-bpel/cart_service/internal/repository"
	"store-bpel/goods_service/schema"
	"testing"
)

var (
	testRepository TestRepository
	testGoods      TestGoodsAdapter
)

// MOCK GOODS ADAPTER
type TestGoodsAdapter interface {
	GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error)
}

type testGoodsAdapter struct {
}

func NewTestGoodsAdapter() TestGoodsAdapter {
	return &testGoodsAdapter{}
}

func (t *testGoodsAdapter) GetProductDetail(ctx context.Context, productId string) (*schema.GetGoodsDefaultResponseData, error) {
	if productId == "invalid-goods" {
		return nil, errors.New("some random error")
	}
	return &schema.GetGoodsDefaultResponseData{
		ListQuantity: []*schema.GetGoodsDefault_QuantityList{
			{
				GoodsSize:  "XL",
				GoodsColor: "red",
				Quantity:   5,
			},
			{
				GoodsSize:  "XL",
				GoodsColor: "yellow",
				Quantity:   1,
			},
			{
				GoodsSize:  "S",
				GoodsColor: "yellow",
				Quantity:   5,
			},
		},
	}, nil
}

// MOCK REPOSITORY
type TestRepository interface {
	AddCart(ctx context.Context, customerId string) error
	GetCart(ctx context.Context, customerId string) (*repository.GetCartModel, error)
	AddGoods(ctx context.Context, cartId string, data []*repository.AddGoodsData) error
	DeleteGoods(ctx context.Context, cartId string, data []*repository.DeleteGoodsData) error
	UpdateGoods(ctx context.Context, cartId string, data []*repository.AddGoodsData) error
	DeleteAllGoods(ctx context.Context, cartId string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) AddCart(ctx context.Context, customerId string) error {
	return nil
}

func (t *testRepo) GetCart(ctx context.Context, customerId string) (*repository.GetCartModel, error) {
	return &repository.GetCartModel{
		CartId: "customer-1",
		Goods: []*repository.GoodsModel{
			{
				CartId:     "cart-1",
				GoodsId:    "goods-1",
				GoodsColor: "red",
				GoodsSize:  "XL",
				Quantity:   2,
			},
			{
				CartId:     "cart-1",
				GoodsId:    "goods-1",
				GoodsColor: "yellow",
				GoodsSize:  "S",
				Quantity:   1,
			},
		},
	}, nil
}

func (t *testRepo) AddGoods(ctx context.Context, cartId string, data []*repository.AddGoodsData) error {
	return nil
}

func (t *testRepo) DeleteGoods(ctx context.Context, cartId string, data []*repository.DeleteGoodsData) error {
	return nil
}

func (t *testRepo) UpdateGoods(ctx context.Context, cartId string, data []*repository.AddGoodsData) error {
	return nil
}

func (t *testRepo) DeleteAllGoods(ctx context.Context, cartId string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testGoods = NewTestGoodsAdapter()

	exitVal := m.Run()

	os.Exit(exitVal)
}
