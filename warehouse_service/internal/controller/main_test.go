package controller

import (
	"context"
	"os"
	"store-bpel/staff_service/schema"
	"store-bpel/warehouse_service/internal/repository"
	"testing"
)

var (
	testRepository TestRepository
	testStaff      TestStaffAdapter
)

// MOCK STAFF ADAPTER
type TestStaffAdapter interface {
	GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error)
}

type testStaffAdapter struct {
}

func NewTestStaffAdapter() TestStaffAdapter {
	return &testStaffAdapter{}
}

func (t *testStaffAdapter) GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error) {
	if staffId == "staff-1" {
		return &schema.GetStaffResponseData{
			StaffId:     "staff-1",
			StaffName:   "Staff One",
			Street:      "THT",
			Ward:        "Ward 11",
			District:    "District 10",
			Province:    "Ho Chi Minh City",
			CitizenId:   "1234567890",
			Role:        "4",
			BranchId:    "branch-1",
			Hometown:    "Ho Chi Minh City",
			Salary:      10000000,
			Birthdate:   "2001-01-01",
			Gender:      "MALE",
			PhoneNumber: "0123456789",
			Status:      "OK",
			Email:       "staff-1@gmail.com",
		}, nil
	}
	return &schema.GetStaffResponseData{
		StaffId:     "staff-2",
		StaffName:   "Staff Two",
		Street:      "THT",
		Ward:        "Ward 11",
		District:    "District 10",
		Province:    "Ho Chi Minh City",
		CitizenId:   "1234567890",
		Role:        "4",
		BranchId:    "branch-1",
		Hometown:    "Ho Chi Minh City",
		Salary:      10000000,
		Birthdate:   "2001-01-01",
		Gender:      "MALE",
		PhoneNumber: "0123456789",
		Status:      "OK",
		Email:       "staff-2@gmail.com",
	}, nil
}

// MOCK REPOSITORY
type TestRepository interface {
	GetStaffByWarehouse(ctx context.Context, warehouseId string) ([]*repository.StaffInWh, error)
	AddWarehouseStaff(ctx context.Context, data *repository.StaffInWh) error
	UpdateWarehouseStaff(ctx context.Context, data *repository.StaffInWh) error
	RemoveWarehouseStaff(ctx context.Context, staffId string) error
	GetWarehouseManager(ctx context.Context, warehouseId string) (*repository.StaffInWh, error)
	UpdateWarehouseManager(ctx context.Context, staffId, warehouseId string) error
	GetWarehouse(ctx context.Context, warehouseId string) (*repository.WarehouseModel, error)
	GetAllWarehouse(ctx context.Context) ([]*repository.WarehouseModel, error)
	AddWarehouse(ctx context.Context, data *repository.WarehouseModel) error
	UpdateWarehouse(ctx context.Context, data *repository.WarehouseModel) error
	DeleteWarehouse(ctx context.Context, warehouseId string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetStaffByWarehouse(ctx context.Context, warehouseId string) ([]*repository.StaffInWh, error) {
	return []*repository.StaffInWh{
		{
			StaffCode: "staff-1",
		},
		{
			StaffCode: "staff-2",
		},
	}, nil
}

func (t *testRepo) AddWarehouseStaff(ctx context.Context, data *repository.StaffInWh) error {
	return nil
}

func (t *testRepo) UpdateWarehouseStaff(ctx context.Context, data *repository.StaffInWh) error {
	return nil
}

func (t *testRepo) RemoveWarehouseStaff(ctx context.Context, staffId string) error {
	return nil
}

func (t *testRepo) GetWarehouseManager(ctx context.Context, warehouseId string) (*repository.StaffInWh, error) {
	return &repository.StaffInWh{
		StaffCode:     "staff-1",
		WarehouseCode: "warehouse-1",
		Role:          "5",
	}, nil
}

func (t *testRepo) UpdateWarehouseManager(ctx context.Context, staffId, warehouseId string) error {
	return nil
}

func (t *testRepo) GetWarehouse(ctx context.Context, warehouseId string) (*repository.WarehouseModel, error) {
	return &repository.WarehouseModel{
		WarehouseCode: "warehouse-1",
		WarehouseName: "Warehouse One",
		Capacity:      1000,
		Street:        "LTK",
		Ward:          "11",
		District:      "10",
		Province:      "HCMC",
	}, nil
}

func (t *testRepo) GetAllWarehouse(ctx context.Context) ([]*repository.WarehouseModel, error) {
	return []*repository.WarehouseModel{
		{
			WarehouseCode: "warehouse-1",
			WarehouseName: "Warehouse One",
			Capacity:      1000,
			Street:        "LTK",
			Ward:          "11",
			District:      "10",
			Province:      "HCMC",
		},
		{
			WarehouseCode: "warehouse-2",
			WarehouseName: "Warehouse Two",
			Capacity:      2000,
			Street:        "THT",
			Ward:          "P. 11",
			District:      "Q. 10",
			Province:      "TP. HCM",
		},
	}, nil
}

func (t *testRepo) AddWarehouse(ctx context.Context, data *repository.WarehouseModel) error {
	return nil
}

func (t *testRepo) UpdateWarehouse(ctx context.Context, data *repository.WarehouseModel) error {
	return nil
}

func (t *testRepo) DeleteWarehouse(ctx context.Context, warehouseId string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testStaff = NewTestStaffAdapter()

	exitVal := m.Run()

	os.Exit(exitVal)
}
