package controller

import (
	"context"
	"store-bpel/warehouse_service/repository"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) UpdateWarehouse(ctx context.Context, request *schema.UpdateWarehouseRequest) error {
	return c.repository.UpdateWarehouse(ctx, &repository.WarehouseModel{
		WarehouseCode: request.WarehouseCode,
		WarehouseName: request.WarehouseName,
		Capacity:      request.Capacity,
		Street:        request.Street,
		Ward:          request.Ward,
		District:      request.District,
		Province:      request.Province,
	})
}
