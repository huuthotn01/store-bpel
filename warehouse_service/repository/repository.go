package repository

import (
	"context"

	"gorm.io/gorm"
)

type IWarehouseServiceRepository interface {
	GetStaffByWarehouse(ctx context.Context, warehouseId string) ([]*StaffInWh, error)
	AddWarehouseStaff(ctx context.Context, data *StaffInWh) error
	UpdateWarehouseStaff(ctx context.Context, data *StaffInWh) error
	RemoveWarehouseStaff(ctx context.Context, staffId string) error
	GetWarehouseManager(ctx context.Context, warehouseId string) (*StaffInWh, error)
	UpdateWarehouseManager(ctx context.Context, staffId, warehouseId string) error
	GetWarehouse(ctx context.Context, warehouseId string) (*WarehouseModel, error)
	GetAllWarehouse(ctx context.Context) ([]*WarehouseModel, error)
	AddWarehouse(ctx context.Context, data *WarehouseModel) error
	UpdateWarehouse(ctx context.Context, data *WarehouseModel) error
	DeleteWarehouse(ctx context.Context, warehouseId string) error
}

func NewRepository(db *gorm.DB) IWarehouseServiceRepository {
	return &warehouseServiceRepository{
		db:                 db,
		warehouseTableName: "warehouse",
		StaffInWhTableName: "staff_in_wh",
	}
}

func (r *warehouseServiceRepository) GetStaffByWarehouse(ctx context.Context, warehouseId string) ([]*StaffInWh, error) {
	var result []*StaffInWh
	query := r.db.WithContext(ctx).Table(r.StaffInWhTableName).Where("warehouse_code = ?", warehouseId)
	return result, query.Find(&result).Error
}

func (r *warehouseServiceRepository) AddWarehouseStaff(ctx context.Context, data *StaffInWh) error {
	return r.db.WithContext(ctx).Table(r.StaffInWhTableName).Create(&data).Error
}

func (r *warehouseServiceRepository) UpdateWarehouseStaff(ctx context.Context, data *StaffInWh) error {
	return r.db.WithContext(ctx).Table(r.StaffInWhTableName).Where("staff_code = ?", data.StaffCode).Select("warehouse_code, role").Updates(&data).Error
}

func (r *warehouseServiceRepository) RemoveWarehouseStaff(ctx context.Context, staffId string) error {
	return r.db.WithContext(ctx).Table(r.StaffInWhTableName).Where("staff_code = ?", staffId).Updates(1).Error
}

func (r *warehouseServiceRepository) GetWarehouseManager(ctx context.Context, warehouseId string) (*StaffInWh, error) {
	var result *StaffInWh
	query := r.db.WithContext(ctx).Table(r.StaffInWhTableName).Where("warehouse_code = ?", warehouseId)
	return result, query.First(&result).Error
}

func (r *warehouseServiceRepository) UpdateWarehouseManager(ctx context.Context, staffId, warehouseId string) error {
	return r.db.WithContext(ctx).Table(r.StaffInWhTableName).Where("warehouse_code = ?", warehouseId).Select("staff_code", "role").Updates(&StaffInWh{
		StaffCode: staffId,
		Role:      "MANAGER",
	}).Error
}

func (r *warehouseServiceRepository) GetWarehouse(ctx context.Context, warehouseId string) (*WarehouseModel, error) {
	var result *WarehouseModel
	query := r.db.WithContext(ctx).Table(r.warehouseTableName).Where("warehouse_code = ?", warehouseId)
	return result, query.First(&result).Error
}

func (r *warehouseServiceRepository) GetAllWarehouse(ctx context.Context) ([]*WarehouseModel, error) {
	var result []*WarehouseModel
	query := r.db.WithContext(ctx).Table(r.warehouseTableName)
	return result, query.Find(&result).Error
}

func (r *warehouseServiceRepository) AddWarehouse(ctx context.Context, data *WarehouseModel) error {
	return r.db.WithContext(ctx).Table(r.warehouseTableName).Create(&data).Error
}

func (r *warehouseServiceRepository) UpdateWarehouse(ctx context.Context, data *WarehouseModel) error {
	return r.db.WithContext(ctx).Table(r.warehouseTableName).Where("warehouse_code = ?", data.WarehouseCode).
		Select("warehouse_code", "warehouse_name", "capacity", "street", "ward", "district", "province").Updates(&data).Error
}

func (r *warehouseServiceRepository) DeleteWarehouse(ctx context.Context, warehouseId string) error {
	return r.db.WithContext(ctx).Table(r.warehouseTableName).Where("warehouse_code = ?", warehouseId).Delete(warehouseId).Error
}
