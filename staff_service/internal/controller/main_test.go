package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	account_schema "store-bpel/account_service/schema"
	"store-bpel/staff_service/internal/repository"
	"testing"
	"time"
)

var (
	testRepository TestRepository
	testAccount    TestAccountAdapter
	testKafka      TestKafkaAdapter
	db             *gorm.DB
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

// MOCK ACCOUNT ADAPTER
type TestAccountAdapter interface {
	UpdateRole(ctx context.Context, username string, request *account_schema.UpdateRoleRequest) error
}

type testAccountAdapter struct {
}

func NewTestAccountAdapter() TestAccountAdapter {
	return &testAccountAdapter{}
}

func (t *testAccountAdapter) UpdateRole(ctx context.Context, username string, request *account_schema.UpdateRoleRequest) error {
	if username == "invalid-account" {
		return errors.New("some random error")
	}
	return nil
}

// MOCK REPOSITORY
type TestRepository interface {
	GetStaff(ctx context.Context, staffName, staffId string) ([]*repository.StaffModel, error)
	AddStaff(ctx context.Context, staff *repository.StaffModel) error
	GetStaffDetail(ctx context.Context, staffId string) (*repository.StaffModel, error)
	UpdateStaff(ctx context.Context, data *repository.StaffModel) error
	DeleteStaffUpdateStatus(ctx context.Context, staffId string) error
	DeleteStaffRemove(ctx context.Context, staffId string) error
	CreateAccount(ctx context.Context, data *repository.AccountModel) error
	GetStaffAttendance(ctx context.Context, staffId string) ([]*repository.AttendanceModel, error)
	CreateStaffRequest(ctx context.Context, request *repository.RequestsModel) error
	UpdateRequestStatus(ctx context.Context, status, requestId string) error
	GetStaffRequest(ctx context.Context, requestId string) (*repository.RequestsModel, error)
	GetListRequest(ctx context.Context) ([]*repository.GetRequestResponseData, error)
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetStaff(ctx context.Context, staffName, staffId string) ([]*repository.StaffModel, error) {
	if staffId == "invalid-staff" {
		return nil, errors.New("some random error")
	}
	return []*repository.StaffModel{
		{
			StaffId:   "staff-1",
			StaffName: "Staff One",
			Gender:    "MALE",
		},
		{
			StaffId:   "staff-2",
			StaffName: "Staff Two",
			Gender:    "FEMALE",
		},
	}, nil
}

func (t *testRepo) AddStaff(ctx context.Context, staff *repository.StaffModel) error {
	if staff.StaffId == "invalid-staff" {
		return errors.New("some random error")
	}
	return nil
}

func (t *testRepo) GetStaffDetail(ctx context.Context, staffId string) (*repository.StaffModel, error) {
	if staffId == "invalid-staff" {
		return nil, errors.New("some random error")
	}
	return &repository.StaffModel{
		StaffId:   "staff-1",
		StaffName: "Staff One",
		Gender:    "MALE",
	}, nil
}

func (t *testRepo) UpdateStaff(ctx context.Context, data *repository.StaffModel) error {
	if data.StaffId == "invalid-staff" {
		return errors.New("some random error")
	}
	return nil
}

func (t *testRepo) DeleteStaffUpdateStatus(ctx context.Context, staffId string) error {
	return nil
}

func (t *testRepo) DeleteStaffRemove(ctx context.Context, staffId string) error {
	return nil
}

func (t *testRepo) CreateAccount(ctx context.Context, data *repository.AccountModel) error {
	if data.StaffId == "invalid-staff-account" {
		return errors.New("some random error")
	}
	return nil
}

func (t *testRepo) GetStaffAttendance(ctx context.Context, staffId string) ([]*repository.AttendanceModel, error) {
	if staffId == "invalid-staff" {
		return nil, errors.New("some random error")
	}
	return []*repository.AttendanceModel{
		{
			StaffId:        "staff-1",
			AttendanceDate: "2023-01-01",
			CheckinTime:    time.Date(2023, 1, 1, 9, 0, 0, 0, time.Local),
			CheckoutTime: sql.NullTime{
				Time: time.Date(2023, 1, 1, 18, 3, 0, 0, time.Local),
			},
		},
		{
			StaffId:        "staff-1",
			AttendanceDate: "2023-01-01",
			CheckinTime:    time.Date(2023, 1, 3, 8, 49, 22, 0, time.Local),
		},
	}, nil
}

func (t *testRepo) CreateStaffRequest(ctx context.Context, request *repository.RequestsModel) error {
	if request.StaffId == "invalid-staff-request" {
		return errors.New("some random error")
	}
	return nil
}

func (t *testRepo) UpdateRequestStatus(ctx context.Context, status, requestId string) error {
	return nil
}

func (t *testRepo) GetStaffRequest(ctx context.Context, requestId string) (*repository.RequestsModel, error) {
	if requestId == "request-1" {
		return &repository.RequestsModel{
			Id:          "request-1",
			RequestType: "ADD",
			StaffId:     "staff-1",
		}, nil
	}
	return &repository.RequestsModel{
		Id:          "request-2",
		RequestType: "DELETE",
		StaffId:     "staff-1",
	}, nil
}

func (t *testRepo) GetListRequest(ctx context.Context) ([]*repository.GetRequestResponseData, error) {
	return []*repository.GetRequestResponseData{
		{
			Id:            "request-1",
			RequestType:   "ADD",
			Status:        "PENDING",
			StaffId:       "staff-2",
			StaffName:     "TNHT",
			Province:      "Binh Duong",
			District:      "Di An",
			Ward:          "Dong Hoa",
			Street:        "Nguyen Du",
			Hometown:      "HCMC",
			CitizenId:     "1234567890",
			StaffPosition: "4",
			Birthdate:     "2001-01-02",
			Salary:        1000000,
			Gender:        "MALE",
			Phone:         "0123456789",
			Email:         "tnht@gmail.com",
			BranchId:      "branch-1",
		},
	}, nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()
	testAccount = NewTestAccountAdapter()
	testKafka = NewTestKafkaAdapter()

	var err error
	db, err = gorm.Open(mysql.Open(dsn()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "bpel", "bpel", "localhost", 3306, "test_db")
}
