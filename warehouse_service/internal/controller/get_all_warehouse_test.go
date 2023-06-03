package controller

import (
	"context"
	"reflect"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_GetAllWarehouse(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		ctx     context.Context
		want    []*schema.GetWarehouseResponseData
		wantErr bool
	}{
		{
			name: "Should get all warehouses correctly",
			ctx:  ctx,
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
		{
			name:    "Should return error when db get all warehouse fails",
			ctx:     context.WithValue(ctx, "status", "fail"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &warehouseServiceController{
				repository: testRepository,
			}
			got, err := c.GetAllWarehouse(tt.ctx)
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
