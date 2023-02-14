package repository

import (
	"context"
	"gorm.io/gorm"
)

type IStaffServiceRepository interface {
	GetStaff(ctx context.Context) ([]*StaffModel, error)
	AddStaff(ctx context.Context, staff *StaffModel, username string) error
	GetStaffDetail(ctx context.Context, staffId string) (*StaffModel, error)
	UpdateStaff(ctx context.Context, data *StaffModel) error
	DeleteStaff(ctx context.Context, staffId string) error
	GetStaffAttendance(ctx context.Context, staffId string) ([]*AttendanceModel, error)
	CreateAddRequest(ctx context.Context, staff *StaffModel, request *RequestsModel, username string) error
	UpdateRequestStatus(ctx context.Context, status, requestId string) error
	DeleteAddRequest(ctx context.Context, staffId string) error
}

func NewRepository(db *gorm.DB) IStaffServiceRepository {
	return &staffServiceRepository{
		db:                  db,
		accountTableName:    "account",
		staffTableName:      "staff",
		attendanceTableName: "attendance",
		requestsTaleName:    "requests",
	}
}

func (r *staffServiceRepository) GetStaff(ctx context.Context) ([]*StaffModel, error) {
	var result []*StaffModel
	query := r.db.WithContext(ctx).Table(r.staffTableName).Find(&result)
	return result, query.Error
}

func (r *staffServiceRepository) GetStaffDetail(ctx context.Context, staffId string) (*StaffModel, error) {
	var result *StaffModel
	query := r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", staffId).Find(&result)
	return result, query.Error
}

func (r *staffServiceRepository) AddStaff(ctx context.Context, staff *StaffModel, username string) error {
	// TODO add to account table in transaction
	return r.db.WithContext(ctx).Table(r.staffTableName).Create(staff).Error
}

func (r *staffServiceRepository) UpdateStaff(ctx context.Context, data *StaffModel) error {
	return nil
}

func (r *staffServiceRepository) DeleteStaff(ctx context.Context, staffId string) error {
	return nil
}

func (r *staffServiceRepository) GetStaffAttendance(ctx context.Context, staffId string) ([]*AttendanceModel, error) {
	return nil, nil
}

func (r *staffServiceRepository) CreateAddRequest(ctx context.Context, staff *StaffModel, request *RequestsModel, username string) error {
	// TODO add new staff and request to both tables
	// TODO add to account table in transaction
	return nil
}

func (r *staffServiceRepository) UpdateRequestStatus(ctx context.Context, status, requestId string) error {
	return nil
}

func (r *staffServiceRepository) DeleteAddRequest(ctx context.Context, staffId string) error {
	return nil
}
