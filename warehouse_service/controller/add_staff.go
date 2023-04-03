package controller

import (
	"context"
	"store-bpel/warehouse_service/repository"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) AddWarehouseStaff(ctx context.Context, request *schema.AddWarehouseStaffRequest) error {
	return c.repository.AddWarehouseStaff(ctx, &repository.StaffInWh{
		StaffCode:     request.StaffId,
		WarehouseCode: request.WarehouseId,
		Role:          request.Role,
	})
}
