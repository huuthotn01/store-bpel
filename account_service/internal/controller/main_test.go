package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"os"
	"store-bpel/account_service/internal/repository"
	"store-bpel/account_service/internal/util"
	customer_schema "store-bpel/customer_service/schema"
	staff_schema "store-bpel/staff_service/schema"
	"testing"
)

var (
	testRepository TestRepository
	testStaff      TestStaffAdapter
	testCustomer   TestCustomerAdapter
	testKafka      TestKafkaAdapter
)

// MOCK REPOSITORY
type TestRepository interface {
	GetListAccount(ctx context.Context, username string) ([]*repository.AccountModel, error)
	GetAccount(ctx context.Context, username string) (*repository.AccountModel, error)
	AddAccount(ctx context.Context, data *repository.AccountModel) error
	UpdateRole(ctx context.Context, username string, role int, password string) error
	UpdatePassword(ctx context.Context, username string, password string) error
	UpdateOTPCode(ctx context.Context, username string, otp string) error
	ConfirmOTP(ctx context.Context, username string, otp string) (*repository.AccountModel, error)
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (r *testRepo) GetListAccount(ctx context.Context, username string) ([]*repository.AccountModel, error) {
	return []*repository.AccountModel{
		{
			UserRole:    1,
			Username:    "cust-1",
			IsActivated: 1,
		},
		{
			UserRole:    3,
			Username:    "staff-1@gmail.com",
			IsActivated: 1,
		},
	}, nil
}

func (r *testRepo) GetAccount(ctx context.Context, username string) (*repository.AccountModel, error) {
	if username == "new-customer" {
		return nil, gorm.ErrRecordNotFound
	}
	pass, err := util.HashPasswordBcrypt("testpwd")
	if err != nil {
		return nil, err
	}
	if username == "unactivated" {
		return &repository.AccountModel{}, nil
	}
	if username == "user-role-7" {
		return &repository.AccountModel{
			UserRole: 7,
			Email:    "user-role-7@gmail.com",
			Username: "user-role-7",
		}, nil
	}
	return &repository.AccountModel{
		Username:    "test-user",
		Email:       "test-user@gmail.com",
		Password:    pass,
		IsActivated: 1,
		UserRole:    1,
	}, nil
}

func (r *testRepo) AddAccount(ctx context.Context, data *repository.AccountModel) error {
	return nil
}

func (r *testRepo) UpdateRole(ctx context.Context, username string, role int, password string) error {
	return nil
}

func (r *testRepo) UpdatePassword(ctx context.Context, username string, password string) error {
	return nil
}

func (r *testRepo) UpdateOTPCode(ctx context.Context, username string, otp string) error {
	return nil
}

func (r *testRepo) ConfirmOTP(ctx context.Context, username string, otp string) (*repository.AccountModel, error) {
	if otp == "invalid-otp" {
		return nil, errors.New("invalid otp")
	}
	return &repository.AccountModel{
		Email: "test-user@gmail.com",
	}, nil
}

// MOCK STAFF ADAPTER
type TestStaffAdapter interface {
	GetDetailStaff(ctx context.Context, staffId string) (*staff_schema.GetStaffResponseData, error)
}

type testStaffAdapter struct {
}

func NewTestStaffAdapter() TestStaffAdapter {
	return &testStaffAdapter{}
}

func (a *testStaffAdapter) GetDetailStaff(ctx context.Context, staffId string) (*staff_schema.GetStaffResponseData, error) {
	return &staff_schema.GetStaffResponseData{
		StaffId:     "staff-1",
		PhoneNumber: "0123456789",
		Email:       "staff-1@gmail.com",
		StaffName:   "Staff One",
	}, nil
}

// MOCK CUSTOMER ADAPTER
type TestCustomerAdapter interface {
	GetCustomer(ctx context.Context, username string) (*customer_schema.GetCustomerInfoResponseData, error)
}

type testCustomerAdapter struct {
}

func NewTestCustomerAdapter() TestCustomerAdapter {
	return &testCustomerAdapter{}
}

func (a *testCustomerAdapter) GetCustomer(ctx context.Context, username string) (*customer_schema.GetCustomerInfoResponseData, error) {
	return &customer_schema.GetCustomerInfoResponseData{
		Phone: "0111111111",
		Email: "cust-1@gmail.com",
		Name:  "Customer One",
	}, nil
}

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

func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testStaff = NewTestStaffAdapter()
	testCustomer = NewTestCustomerAdapter()
	testKafka = NewTestKafkaAdapter()

	exitVal := m.Run()

	os.Exit(exitVal)
}
