package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) GetWarehouse(ctx context.Context, warehouseId string) (*schema.GetWarehouseResponseData, error) {
	warehouseData, err := c.repository.GetWarehouse(ctx, warehouseId)
	if err != nil {
		return nil, err
	}

	return &schema.GetWarehouseResponseData{
		WarehouseName: warehouseData.WarehouseName,
		Capacity:      warehouseData.Capacity,
		CreatedAt:     warehouseData.CreatedAt,
		Street:        warehouseData.Street,
		Ward:          warehouseData.Ward,
		District:      warehouseData.District,
		Province:      warehouseData.Province,
	}, nil
}
