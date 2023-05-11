package controller

import (
	"context"
	"reflect"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_GetStaff(t *testing.T) {
	type args struct {
		staffName string
		staffId   string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetStaffResponseData
		wantErr bool
	}{
		{
			name: "Should get staff correctly",
			args: args{
				staffId: "staff-1",
			},
			want: []*schema.GetStaffResponseData{
				{
					StaffId:   "staff-1",
					StaffName: "Staff One",
					Gender:    "MALE",
				},
				{
					StaffId:   "staff-2",
					StaffName: "Staff Two",
					Gender:    "FEMALE",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				repository: testRepository,
			}
			got, err := s.GetStaff(ctx, tt.args.staffName, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_staffServiceController_GetDetailStaff(t *testing.T) {
	type args struct {
		staffId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetStaffResponseData
		wantErr bool
	}{
		{
			name: "Should get detail staff correctly",
			args: args{
				staffId: "staff-1",
			},
			want: &schema.GetStaffResponseData{
				StaffId:   "staff-1",
				StaffName: "Staff One",
				Gender:    "MALE",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				repository: testRepository,
			}
			got, err := s.GetDetailStaff(ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetailStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDetailStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
