package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
)

type IWarehouseBffController interface {
	GetManager(ctx context.Context, request *warehouse_service.GetWarehouseId) (*warehouse_service.GetWarehouseManagerResponseData, error)
	UpdateManager(ctx context.Context, request *warehouse_service.UpdateManagerRequest) error
	GetStaff(ctx context.Context, request *warehouse_service.GetWarehouseId) ([]*warehouse_service.GetWarehouseStaffResponseData, error)
	AddStaff(ctx context.Context, request *warehouse_service.AddWarehouseStaffRequest) error
	UpdateStaff(ctx context.Context, request *warehouse_service.UpdateWarehouseStaffRequest) error
	DeleteStaff(ctx context.Context, request *warehouse_service.DeleteWarehouseStaffRequest) error
	GetWarehouse(ctx context.Context, request *warehouse_service.GetWarehouseId) (*warehouse_service.GetWarehouseResponseData, error)
	AddWarehouse(ctx context.Context, request *warehouse_service.AddWarehouseRequest) error
	UpdateWarehouse(ctx context.Context, request *warehouse_service.UpdateWarehouseRequest) error
	DeleteWarehouse(ctx context.Context, request *warehouse_service.DeleteWarehouseRequest) error
}

type warehouseBffController struct {
	cfg              *config.Config
	warehouseAdapter adapter.IWarehouseServiceAdapter
}

func NewController(cfg *config.Config) IWarehouseBffController {
	// init warehouse adapter
	warehouseAdapter := adapter.NewWarehouseAdapter(cfg)

	return &warehouseBffController{
		cfg:              cfg,
		warehouseAdapter: warehouseAdapter,
	}
}
