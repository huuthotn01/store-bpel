package controller

import (
	"context"
	"reflect"
	"store-bpel/warehouse_service/config"
	"store-bpel/warehouse_service/schema"
	"testing"
)

func Test_warehouseServiceController_GetWarehouseStaff(t *testing.T) {
	type args struct {
		warehouseId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetWarehouseStaffResponseData
		wantErr bool
	}{
		{
			name: "Should get warehouse staff correctly",
			args: args{
				warehouseId: "warehouse-1",
			},
			want: []*schema.GetWarehouseStaffResponseData{
				{
					StaffId:     "staff-1",
					StaffName:   "Staff One",
					Street:      "THT",
					Ward:        "Ward 11",
					District:    "District 10",
					Province:    "Ho Chi Minh City",
					CitizenId:   "1234567890",
					Role:        "4",
					BranchId:    "branch-1",
					Hometown:    "Ho Chi Minh City",
					Salary:      10000000,
					Birthdate:   "2001-01-01",
					Gender:      "MALE",
					PhoneNumber: "0123456789",
					Status:      "OK",
					Email:       "staff-1@gmail.com",
				},
				{
					StaffId:     "staff-2",
					StaffName:   "Staff Two",
					Street:      "THT",
					Ward:        "Ward 11",
					District:    "District 10",
					Province:    "Ho Chi Minh City",
					CitizenId:   "1234567890",
					Role:        "4",
					BranchId:    "branch-1",
					Hometown:    "Ho Chi Minh City",
					Salary:      10000000,
					Birthdate:   "2001-01-01",
					Gender:      "MALE",
					PhoneNumber: "0123456789",
					Status:      "OK",
					Email:       "staff-2@gmail.com",
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
				config:       cfg,
				repository:   testRepository,
				staffAdapter: testStaff,
			}
			got, err := c.GetWarehouseStaff(ctx, tt.args.warehouseId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWarehouseStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWarehouseStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
