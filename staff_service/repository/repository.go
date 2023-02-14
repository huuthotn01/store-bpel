package repository

import (
	"context"
	"gorm.io/gorm"
)

type IStaffServiceRepository interface {
	GetStaff(ctx context.Context) ([]*StaffModel, error)
	AddStaff(ctx context.Context, staff *StaffModel) error
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

func (r *staffServiceRepository) AddStaff(ctx context.Context, staff *StaffModel) error {
	return r.db.WithContext(ctx).Table(r.staffTableName).Create(staff).Error
}
