package controller

import (
	"context"
	"fmt"
	"store-bpel/warehouse_service/repository"
	"store-bpel/warehouse_service/schema"
	"time"
)

func (c *warehouseServiceController) AddWarehouse(ctx context.Context, request *schema.AddWarehouseRequest) error {
	warehouseCode := fmt.Sprintf("warehouse-%s", time.Now().UTC().String())

	return c.repository.AddWarehouse(ctx, &repository.WarehouseModel{
		WarehouseCode: warehouseCode,
		WarehouseName: request.WarehouseName,
		Capacity:      request.Capacity,
		Street:        request.Street,
		Ward:          request.Ward,
		District:      request.District,
		Province:      request.Province,
	})
}
