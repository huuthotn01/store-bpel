package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) DeleteStaff(ctx context.Context, request *schema.DeleteWarehouseStaffRequest) error {
	return c.repository.RemoveWarehouseStaff(ctx, request.StaffId)
}
