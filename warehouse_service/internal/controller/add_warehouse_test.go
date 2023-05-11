package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_AddWarehouse(t *testing.T) {
	type args struct {
		request *schema.AddWarehouseRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add warehouse successfully",
			args: args{
				request: &schema.AddWarehouseRequest{
					WarehouseName: "New Warehouse",
					Capacity:      1000,
					Street:        "To Hien Thanh",
					Ward:          "Ward 11",
					District:      "District 10",
					Province:      "Ho Chi Minh City",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &warehouseServiceController{
				repository: testRepository,
			}
			if err := c.AddWarehouse(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddWarehouse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
