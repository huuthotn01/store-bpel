package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) UpdateStaff(ctx context.Context, request *warehouse_service.UpdateWarehouseStaffRequest) error {
	return c.warehouseAdapter.UpdateStaff(ctx, &schema.UpdateWarehouseStaffRequest{
		StaffId:     request.StaffId,
		WarehouseId: request.WarehouseId,
		Role:        request.Role,
	})
}
