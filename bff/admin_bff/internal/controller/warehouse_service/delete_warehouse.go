package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) DeleteWarehouse(ctx context.Context, request *warehouse_service.DeleteWarehouseRequest) error {
	return c.warehouseAdapter.DeleteWarehouse(ctx, &schema.DeleteWarehouseRequest{
		WarehouseCode: request.WarehouseCode,
	})
}
