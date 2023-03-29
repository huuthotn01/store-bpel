package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/warehouse_service/adapter"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
)

type IWarehouseServiceController interface {
	GetWarehouseStaff(ctx context.Context) ([]*schema.GetWarehouseStaffResponseData, error)
}

type warehouseServiceController struct {
	config     *config.Config
	repository *gorm.DB

	staffAdapter adapter.IStaffServiceAdapter
}

func NewController(config *config.Config, db *gorm.DB) IWarehouseServiceController {
	// init staff adapter
	staffAdapter := adapter.NewStaffAdapter(config)

	return &warehouseServiceController{
		config:       config,
		repository:   db,
		staffAdapter: staffAdapter,
	}
}
