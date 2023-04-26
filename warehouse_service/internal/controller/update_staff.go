package controller

import (
	"context"
	"store-bpel/warehouse_service/internal/repository"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) UpdateStaff(ctx context.Context, request *schema.UpdateWarehouseStaffRequest) error {
	return c.repository.UpdateWarehouseStaff(ctx, &repository.StaffInWh{
		StaffCode:     request.StaffId,
		WarehouseCode: request.WarehouseId,
		Role:          request.Role,
	})
}
