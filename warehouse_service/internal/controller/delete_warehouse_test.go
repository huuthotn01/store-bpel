package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_DeleteWarehouse(t *testing.T) {
	type args struct {
		request *schema.DeleteWarehouseRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete warehouse successfully",
			args: args{
				request: &schema.DeleteWarehouseRequest{
					WarehouseCode: "warehouse-1",
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
			if err := c.DeleteWarehouse(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("DeleteWarehouse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
