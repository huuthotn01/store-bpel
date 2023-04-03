package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) UpdateWarehouse(ctx context.Context, request *warehouse_service.UpdateWarehouseRequest) error {
	return c.warehouseAdapter.UpdateWarehouse(ctx, &schema.UpdateWarehouseRequest{
		WarehouseCode: request.WarehouseCode,
		WarehouseName: request.WarehouseName,
		Capacity:      request.Capacity,
		Street:        request.Street,
		Ward:          request.Ward,
		District:      request.District,
		Province:      request.Province,
	})
}
