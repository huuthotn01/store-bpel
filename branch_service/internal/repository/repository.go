package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBranchServiceRepository interface {
	GetBranch(ctx context.Context) ([]*BranchModel, error)
	GetBranchDetail(ctx context.Context, branchId string) (*BranchModel, error)
	GetBranchStaff(ctx context.Context, branchId string) ([]*BranchStaffModel, error)
	AddBranchStaff(ctx context.Context, data *BranchStaffModel) error
	AddBranch(ctx context.Context, data *BranchModel) error
	UpdateBranch(ctx context.Context, data *BranchModel) error
	UpdateBranchManager(ctx context.Context, branchId string, managerId string) error
	DeleteBranch(ctx context.Context, branchId string) error
}

func NewRepository(db *gorm.DB) IBranchServiceRepository {
	return &branchServiceRepository{
		db:                     db,
		branchTableName:        "branch",
		branchImgTableName:     "branch_img",
		branchManagerTableName: "branch_manager",
		branchStaffTableName:   "branch_staff",
	}
}

func (r *branchServiceRepository) GetBranch(ctx context.Context) ([]*BranchModel, error) {
	var result []*BranchModel
	query := r.db.WithContext(ctx).Table(r.branchTableName).Find(&result)
	return result, query.Error
}

func (r *branchServiceRepository) GetBranchDetail(ctx context.Context, branchId string) (*BranchModel, error) {
	var result *BranchModel
	query := r.db.WithContext(ctx).Table(r.branchTableName).Where("branch_code = ?", branchId).First(&result)
	return result, query.Error
}

func (r *branchServiceRepository) GetBranchStaff(ctx context.Context, branchId string) ([]*BranchStaffModel, error) {
	result := make([]*BranchStaffModel, 0)
	query := r.db.WithContext(ctx).Table(r.branchStaffTableName).Where("branch_code = ?", branchId).Find(&result)
	return result, query.Error
}

func (r *branchServiceRepository) AddBranchStaff(ctx context.Context, data *BranchStaffModel) error {
	// check if working place sent by staff exists
	_, err := r.GetBranchDetail(ctx, data.BranchCode)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Table(r.branchStaffTableName).Select("branch_code", "staff_code").Create(&data).Error
}

func (r *branchServiceRepository) AddBranch(ctx context.Context, data *BranchModel) error {
	return r.db.WithContext(ctx).Table(r.branchTableName).Create(data).Error
}

func (r *branchServiceRepository) UpdateBranch(ctx context.Context, data *BranchModel) error {
	return r.db.WithContext(ctx).Table(r.branchTableName).Where("branch_code = ?", data.BranchCode).Updates(data).Error
}

func (r *branchServiceRepository) UpdateBranchManager(ctx context.Context, branchId string, managerId string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// end current manager
		if err := r.db.WithContext(ctx).Exec("update branch_manager set end_date = NOW() where branch_code = ? and end_date is null", branchId).Error; err != nil {
			return err
		}
		// add new manager
		if err := r.db.WithContext(ctx).Table(r.branchManagerTableName).Select("branch_code", "manager_code").Create(&BranchManagerModel{
			BranchCode:  branchId,
			ManagerCode: managerId,
		}).Error; err != nil {
			return err
		}
		if err := r.db.WithContext(ctx).Table(r.branchStaffTableName).Clauses(
			clause.OnConflict{
				DoNothing: true,
			},
		).Select("branch_code", "staff_code").Create(&BranchStaffModel{
			BranchCode: branchId,
			StaffCode:  managerId,
		}).Error; err != nil {
			return err
		}
		if err := r.db.WithContext(ctx).Table(r.branchTableName).Where("branch_code = ?", branchId).Update("manager", managerId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *branchServiceRepository) DeleteBranch(ctx context.Context, branchId string) error {
	return r.db.WithContext(ctx).Table(r.branchTableName).Where("branch_code = ?", branchId).Delete(branchId).Error
}
