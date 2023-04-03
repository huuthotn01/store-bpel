package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) UpdateManager(ctx context.Context, request *warehouse_service.UpdateManagerRequest) error {
	return c.warehouseAdapter.UpdateManager(ctx, &schema.UpdateManagerRequest{
		StaffId:     request.StaffId,
		WarehouseId: request.WarehouseId,
	})
}
