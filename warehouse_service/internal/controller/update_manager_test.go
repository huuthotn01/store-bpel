package controller

import (
	"context"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_UpdateWarehouseManager(t *testing.T) {
	type args struct {
		request *schema.UpdateManagerRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update warehouse manager successfully",
			args: args{
				request: &schema.UpdateManagerRequest{
					StaffId:     "staff-1",
					WarehouseId: "warehouse-1",
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
			if err := c.UpdateWarehouseManager(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateWarehouseManager() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
