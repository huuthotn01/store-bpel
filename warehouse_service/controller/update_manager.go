package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) UpdateWarehouseManager(ctx context.Context, request *schema.UpdateManagerRequest) error {
	return c.repository.UpdateWarehouseManager(ctx, request.StaffId, request.WarehouseId)
}
