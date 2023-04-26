package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseBffController) AddWarehouse(ctx context.Context, request *warehouse_service.AddWarehouseRequest) error {
	return c.warehouseAdapter.AddWarehouse(ctx, &schema.AddWarehouseRequest{
		WarehouseName: request.WarehouseName,
		Capacity:      request.Capacity,
		Street:        request.Street,
		Ward:          request.Ward,
		District:      request.District,
		Province:      request.Province,
	})
}
