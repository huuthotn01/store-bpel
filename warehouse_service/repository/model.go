package repository

import (
	"gorm.io/gorm"
	"time"
)

type warehouseServiceRepository struct {
	db                 *gorm.DB
	warehouseTableName string
	StaffInWhTableName string
}

type WarehouseModel struct {
	WarehouseCode string
	WarehouseName string
	Capacity      int
	CreatedAt     time.Time
	Street        string
	Ward          string
	District      string
	Province      string
}

type StaffInWh struct {
	StaffCode     string
	WarehouseCode string
	StartedDate   time.Time
	EndDate       time.Time
	Role          string
}
