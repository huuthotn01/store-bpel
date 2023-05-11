package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_UpdateStaff(t *testing.T) {
	type args struct {
		request *schema.UpdateWarehouseStaffRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update staff successfully",
			args: args{
				request: &schema.UpdateWarehouseStaffRequest{
					StaffId:     "staff-1",
					WarehouseId: "warehouse-1",
					Role:        "4",
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
			if err := c.UpdateStaff(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
