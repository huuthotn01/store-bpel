package controller

import (
	"context"
	"gorm.io/gorm"
	"os"
	"store-bpel/branch_service/internal/repository"
	"testing"
	"time"
)

var (
	testRepository TestRepository
)

// MOCK REPOSITORY
type TestRepository interface {
	GetBranch(ctx context.Context) ([]*repository.BranchModel, error)
	GetBranchDetail(ctx context.Context, branchId string) (*repository.BranchModel, error)
	GetBranchStaff(ctx context.Context, branchId string) ([]*repository.BranchStaffModel, error)
	AddBranchStaff(ctx context.Context, data *repository.BranchStaffModel) error
	AddBranch(ctx context.Context, data *repository.BranchModel) error
	UpdateBranch(ctx context.Context, data *repository.BranchModel) error
	UpdateBranchManager(ctx context.Context, branchId string, managerId string) error
	DeleteBranch(ctx context.Context, branchId string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetBranch(ctx context.Context) ([]*repository.BranchModel, error) {
	return []*repository.BranchModel{
		{
			BranchCode:     "branch-1",
			BranchName:     "Branch One",
			BranchStreet:   "LTK",
			BranchWard:     "Ward 11",
			BranchDistrict: "District 10",
			BranchProvince: "HCMC",
			CreatedAt:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.Local),
			Manager:        "staff-1",
			OpenTime:       "07:00",
			CloseTime:      "21:00",
		},
		{
			BranchCode:     "branch-2",
			BranchName:     "Branch Two",
			BranchStreet:   "THT",
			BranchWard:     "P.11",
			BranchDistrict: "Q.10",
			BranchProvince: "TP HCM",
			CreatedAt:      time.Date(2022, 02, 02, 0, 0, 0, 0, time.Local),
			Manager:        "staff-2",
			OpenTime:       "09:00",
			CloseTime:      "18:00",
		},
	}, nil
}

func (t *testRepo) GetBranchDetail(ctx context.Context, branchId string) (*repository.BranchModel, error) {
	if branchId == "branch-1" {
		return &repository.BranchModel{
			BranchCode:     "branch-1",
			BranchName:     "Branch One",
			BranchStreet:   "LTK",
			BranchWard:     "Ward 11",
			BranchDistrict: "District 10",
			BranchProvince: "HCMC",
			CreatedAt:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.Local),
			Manager:        "staff-1",
			OpenTime:       "07:00",
			CloseTime:      "21:00",
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (t *testRepo) GetBranchStaff(ctx context.Context, branchId string) ([]*repository.BranchStaffModel, error) {
	return []*repository.BranchStaffModel{
		{
			StaffCode: "staff-1",
		},
		{
			StaffCode: "staff-2",
		},
		{
			StaffCode: "staff-3",
		},
	}, nil
}

func (t *testRepo) AddBranchStaff(ctx context.Context, data *repository.BranchStaffModel) error {
	return nil
}

func (t *testRepo) AddBranch(ctx context.Context, data *repository.BranchModel) error {
	return nil
}

func (t *testRepo) UpdateBranch(ctx context.Context, data *repository.BranchModel) error {
	return nil
}

func (t *testRepo) UpdateBranchManager(ctx context.Context, branchId string, managerId string) error {
	return nil
}

func (t *testRepo) DeleteBranch(ctx context.Context, branchId string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()

	exitVal := m.Run()

	os.Exit(exitVal)
}
