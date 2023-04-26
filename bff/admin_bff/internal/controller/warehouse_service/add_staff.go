package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) AddStaff(ctx context.Context, request *warehouse_service.AddWarehouseStaffRequest) error {
	return c.warehouseAdapter.AddStaff(ctx, &schema.AddWarehouseStaffRequest{
		StaffId:     request.StaffId,
		WarehouseId: request.WarehouseId,
		Role:        request.Role,
	})
}
