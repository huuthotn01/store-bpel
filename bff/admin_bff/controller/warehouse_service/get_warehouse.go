package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
)

func (c *warehouseBffController) GetWarehouse(ctx context.Context, request *warehouse_service.GetWarehouseId) (*warehouse_service.GetWarehouseResponseData, error) {
	warehouse, err := c.warehouseAdapter.GetWarehouse(ctx, request.WarehouseId)
	if err != nil {
		return nil, err
	}

	return &warehouse_service.GetWarehouseResponseData{
		WarehouseCode: warehouse.WarehouseCode,
		WarehouseName: warehouse.WarehouseName,
		Capacity:      warehouse.Capacity,
		CreatedAt:     warehouse.CreatedAt,
		Street:        warehouse.Street,
		Ward:          warehouse.Ward,
		District:      warehouse.District,
		Province:      warehouse.Province,
	}, nil
}
