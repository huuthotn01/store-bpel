package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) DeleteStaff(ctx context.Context, request *warehouse_service.DeleteWarehouseStaffRequest) error {
	return c.warehouseAdapter.DeleteStaff(ctx, &schema.DeleteWarehouseStaffRequest{
		StaffId: request.StaffId,
	})
}
