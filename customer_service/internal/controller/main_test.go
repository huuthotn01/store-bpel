package controller

import (
	"context"
	"gorm.io/gorm"
	"os"
	"store-bpel/customer_service/internal/repository"
	"testing"
)

var (
	testRepository TestRepository
)

// MOCK REPOSITORY
type TestRepository interface {
	GetCustomerInfo(ctx context.Context, customerId string) (*repository.CustomerModel, error)
	UpdateCustomerInfo(ctx context.Context, data *repository.CustomerModel) error
	AddCustomer(ctx context.Context, data *repository.CustomerModel) error
	UpdateCustomerImage(ctx context.Context, customerId, imageUrl string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetCustomerInfo(ctx context.Context, customerId string) (*repository.CustomerModel, error) {
	if customerId == "huutho" {
		return nil, gorm.ErrRecordNotFound
	}
	return &repository.CustomerModel{
		Username:       "httn",
		CustomerEmail:  "httn@gmail.com",
		CustomerName:   "Huu Tho",
		CustomerPhone:  "0111111111",
		CustomerGender: "MALE",
		Street:         "THT",
		Ward:           "11",
		District:       "10",
		Province:       "HCMC",
	}, nil
}

func (t *testRepo) UpdateCustomerInfo(ctx context.Context, data *repository.CustomerModel) error {
	return nil
}

func (t *testRepo) AddCustomer(ctx context.Context, data *repository.CustomerModel) error {
	return nil
}

func (t *testRepo) UpdateCustomerImage(ctx context.Context, customerId, imageUrl string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()

	exitVal := m.Run()

	os.Exit(exitVal)
}
