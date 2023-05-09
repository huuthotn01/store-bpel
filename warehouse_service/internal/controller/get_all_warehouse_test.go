package controller

import (
	"context"
	"reflect"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_GetAllWarehouse(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetWarehouseResponseData
		wantErr bool
	}{
		{
			name: "Should get all warehouses correctly",
			want: []*schema.GetWarehouseResponseData{
				{
					WarehouseCode: "warehouse-1",
					WarehouseName: "Warehouse One",
					Capacity:      1000,
					Street:        "LTK",
					Ward:          "11",
					District:      "10",
					Province:      "HCMC",
				},
				{
					WarehouseCode: "warehouse-2",
					WarehouseName: "Warehouse Two",
					Capacity:      2000,
					Street:        "THT",
					Ward:          "P. 11",
					District:      "Q. 10",
					Province:      "TP. HCM",
				},
			},
		},
	}

	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &warehouseServiceController{
				config:     cfg,
				repository: testRepository,
			}
			got, err := c.GetAllWarehouse(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllWarehouse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
