package controller

import (
	"context"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_UpdateWarehouse(t *testing.T) {
	type args struct {
		request *schema.UpdateWarehouseRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update warehouse successfully",
			args: args{
				request: &schema.UpdateWarehouseRequest{
					WarehouseCode: "warehouse-1",
					WarehouseName: "New Warehouse Name",
					Capacity:      5000,
					Street:        "THT",
					Ward:          "Ward 10",
					District:      "Dist 10",
					Province:      "HCMC",
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
			if err := c.UpdateWarehouse(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateWarehouse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
