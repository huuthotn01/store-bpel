package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) DeleteWarehouse(ctx context.Context, request *schema.DeleteWarehouseRequest) error {
	return c.repository.DeleteWarehouse(ctx, request.WarehouseCode)
}
