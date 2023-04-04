package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
)

func (c *warehouseBffController) GetAllWarehouse(ctx context.Context) ([]*warehouse_service.GetWarehouseResponseData, error) {
	warehouses, err := c.warehouseAdapter.GetAllWarehouse(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*warehouse_service.GetWarehouseResponseData, 0, len(warehouses))
	for _, warehouse := range warehouses {
		result = append(result, &warehouse_service.GetWarehouseResponseData{
			WarehouseCode: warehouse.WarehouseCode,
			WarehouseName: warehouse.WarehouseName,
			Capacity:      warehouse.Capacity,
			CreatedAt:     warehouse.CreatedAt,
			Street:        warehouse.Street,
			Ward:          warehouse.Ward,
			District:      warehouse.District,
			Province:      warehouse.Province,
		})
	}

	return result, nil
}
