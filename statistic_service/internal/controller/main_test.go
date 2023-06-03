package controller

import (
	"context"
	"errors"
	"os"
	"store-bpel/statistic_service/internal/repository"
	"testing"
)

var (
	testRepository TestRepository
)

// MOCK REPOSITORY
type TestRepository interface {
	AddOrderData(ctx context.Context, goodsModel []*repository.GoodsModel, orderModel *repository.OrdersModel) error
	GetOverallStat(ctx context.Context, start, end string, goodsId string, branchId []string, gender []int, goodsType []string) ([]*repository.OverallStatData, error)
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) AddOrderData(ctx context.Context, goodsModel []*repository.GoodsModel, orderModel *repository.OrdersModel) error {
	return nil
}

func (t *testRepo) GetOverallStat(ctx context.Context, start, end string, goodsId string, branchId []string, gender []int, goodsType []string) ([]*repository.OverallStatData, error) {
	if len(branchId) > 0 && branchId[0] == "invalid-branch" {
		return nil, errors.New("some random error")
	}
	if goodsId == "invalid-goods" {
		return nil, errors.New("some random error")
	}
	return []*repository.OverallStatData{
		{
			Revenue: 1000,
			Profit:  500,
			Date:    "2023-01-01",
		},
		{
			Revenue: 2000,
			Profit:  600,
			Date:    "2023-01-02",
		},
		{
			Revenue: 100,
			Profit:  10,
			Date:    "2023-01-03",
		},
	}, nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()

	exitVal := m.Run()

	os.Exit(exitVal)
}
