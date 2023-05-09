package controller

import (
	"context"
	"reflect"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_GetWarehouse(t *testing.T) {
	type args struct {
		warehouseId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetWarehouseResponseData
		wantErr bool
	}{
		{
			name: "Should get warehouse successfully",
			args: args{
				warehouseId: "warehouse-1",
			},
			want: &schema.GetWarehouseResponseData{
				WarehouseCode: "warehouse-1",
				WarehouseName: "Warehouse One",
				Capacity:      1000,
				Street:        "LTK",
				Ward:          "11",
				District:      "10",
				Province:      "HCMC",
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
			got, err := c.GetWarehouse(ctx, tt.args.warehouseId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWarehouse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
