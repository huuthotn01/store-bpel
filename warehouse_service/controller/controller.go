package controller

import (
	"context"
	"store-bpel/warehouse_service/adapter"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/repository"
	"store-bpel/warehouse_service/schema"

	"gorm.io/gorm"
)

type IWarehouseServiceController interface {
	GetWarehouseManager(ctx context.Context, warehouseId string) (*schema.GetWarehouseManagerResponseData, error)
	UpdateWarehouseManager(ctx context.Context, request *schema.UpdateManagerRequest) error
	GetWarehouseStaff(ctx context.Context, warehouseId string) ([]*schema.GetWarehouseStaffResponseData, error)
	AddWarehouseStaff(ctx context.Context, request *schema.AddWarehouseStaffRequest) error
	UpdateStaff(ctx context.Context, request *schema.UpdateWarehouseStaffRequest) error
	DeleteStaff(ctx context.Context, request *schema.DeleteWarehouseStaffRequest) error
	GetWarehouse(ctx context.Context, warehouseId string) (*schema.GetWarehouseResponseData, error)
	GetAllWarehouse(ctx context.Context) ([]*schema.GetWarehouseResponseData, error)
	AddWarehouse(ctx context.Context, request *schema.AddWarehouseRequest) error
	UpdateWarehouse(ctx context.Context, request *schema.UpdateWarehouseRequest) error
	DeleteWarehouse(ctx context.Context, request *schema.DeleteWarehouseRequest) error
}

type warehouseServiceController struct {
	config     *config.Config
	repository repository.IWarehouseServiceRepository

	staffAdapter adapter.IStaffServiceAdapter
}

func NewController(config *config.Config, db *gorm.DB) IWarehouseServiceController {
	// init staff adapter
	staffAdapter := adapter.NewStaffAdapter(config)

	// init repository
	repo := repository.NewRepository(db)

	return &warehouseServiceController{
		config:       config,
		repository:   repo,
		staffAdapter: staffAdapter,
	}
}
