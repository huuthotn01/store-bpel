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
	CreateStaffRequest(ctx context.Context, request *RequestsModel) error
	UpdateRequestStatus(ctx context.Context, status, requestId string) error
	GetStaffRequest(ctx context.Context, requestId string) (*RequestsModel, error)
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		// add to staff table
		err := tx.WithContext(ctx).Table(r.staffTableName).Create(staff).Error
		if err != nil {
			return err
		}
		// add to account table
		return tx.WithContext(ctx).Table(r.accountTableName).Create(&AccountModel{
			Username: username,
			StaffId:  staff.StaffId,
		}).Error
	})
}

func (r *staffServiceRepository) UpdateStaff(ctx context.Context, data *StaffModel) error {
	return r.db.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", data.StaffId).Updates(data).Error
}

func (r *staffServiceRepository) DeleteStaff(ctx context.Context, staffId string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// delete in account table
		err := tx.WithContext(ctx).Table(r.accountTableName).Where("staff_id = ?", staffId).Delete(staffId).Error
		if err != nil {
			return err
		}
		// delete in staff table
		return tx.WithContext(ctx).Table(r.staffTableName).Where("staff_id = ?", staffId).Delete(staffId).Error
	})
}

func (r *staffServiceRepository) GetStaffAttendance(ctx context.Context, staffId string) ([]*AttendanceModel, error) {
	result := make([]*AttendanceModel, 0)
	query := r.db.WithContext(ctx).Table(r.attendanceTableName).Where("staff_id = ?", staffId).Find(&result)
	return result, query.Error
}

func (r *staffServiceRepository) CreateStaffRequest(ctx context.Context, request *RequestsModel) error {
	return r.db.WithContext(ctx).Table(r.requestsTaleName).Create(request).Error
}

func (r *staffServiceRepository) UpdateRequestStatus(ctx context.Context, status, requestId string) error {
	return r.db.WithContext(ctx).Table(r.requestsTaleName).Where("id = ?", requestId).Update("status", status).Error
}

func (r *staffServiceRepository) GetStaffRequest(ctx context.Context, requestId string) (*RequestsModel, error) {
	var result *RequestsModel
	query := r.db.WithContext(ctx).Table(r.requestsTaleName).Where("id = ?", requestId).First(&result)
	return result, query.Error
}
