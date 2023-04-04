package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) GetAllWarehouse(ctx context.Context) ([]*schema.GetWarehouseResponseData, error) {
	warehouseDatas, err := c.repository.GetAllWarehouse(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*schema.GetWarehouseResponseData, 0, len(warehouseDatas))
	for _, warehouse := range warehouseDatas {
		result = append(result, &schema.GetWarehouseResponseData{
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
