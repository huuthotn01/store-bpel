package controller

import (
	"context"
	"reflect"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_GetWarehouseManager(t *testing.T) {
	type args struct {
		warehouseId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetWarehouseManagerResponseData
		wantErr bool
	}{
		{
			name: "Should get warehouse manager correctly",
			args: args{
				warehouseId: "warehouse-1",
			},
			want: &schema.GetWarehouseManagerResponseData{
				StaffId:     "staff-1",
				StaffName:   "Staff One",
				Street:      "THT",
				Ward:        "Ward 11",
				District:    "District 10",
				Province:    "Ho Chi Minh City",
				CitizenId:   "1234567890",
				BranchId:    "branch-1",
				Hometown:    "Ho Chi Minh City",
				Salary:      10000000,
				Birthdate:   "2001-01-01",
				Gender:      "MALE",
				PhoneNumber: "0123456789",
				Status:      "OK",
				Email:       "staff-1@gmail.com",
			},
		},
		{
			name: "Should return error when db get warehouse manager fails",
			args: args{
				warehouseId: "invalid-warehouse",
			},
			wantErr: true,
		},
		{
			name: "Should return error when staff adapter fails",
			args: args{
				warehouseId: "invalid-warehouse-manager",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &warehouseServiceController{
				repository:   testRepository,
				staffAdapter: testStaff,
			}
			got, err := c.GetWarehouseManager(ctx, tt.args.warehouseId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWarehouseManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWarehouseManager() got = %v, want %v", got, tt.want)
			}
		})
	}
}
