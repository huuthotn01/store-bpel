package controller

import (
	"context"
	"reflect"
	"store-bpel/branch_service/schema"
	"testing"
	"time"
)

func Test_branchServiceController_GetBranch(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetBranchResponseData
		wantErr bool
	}{
		{
			name: "Should get branches correctly",
			want: []*schema.GetBranchResponseData{
				{
					BranchCode:     "branch-1",
					BranchName:     "Branch One",
					BranchStreet:   "LTK",
					BranchWard:     "Ward 11",
					BranchDistrict: "District 10",
					BranchProvince: "HCMC",
					CreatedAt:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.Local),
					Manager:        "staff-1",
					OpenTime:       "07:00",
					CloseTime:      "21:00",
				},
				{
					BranchCode:     "branch-2",
					BranchName:     "Branch Two",
					BranchStreet:   "THT",
					BranchWard:     "P.11",
					BranchDistrict: "Q.10",
					BranchProvince: "TP HCM",
					CreatedAt:      time.Date(2022, 02, 02, 0, 0, 0, 0, time.Local),
					Manager:        "staff-2",
					OpenTime:       "09:00",
					CloseTime:      "18:00",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			got, err := s.GetBranch(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_branchServiceController_GetBranchDetail(t *testing.T) {
	type args struct {
		branchId string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.GetBranchResponseData
		wantErr bool
	}{
		{
			name: "Should get branch correctly",
			args: args{
				branchId: "branch-1",
			},
			want: &schema.GetBranchResponseData{
				BranchCode:     "branch-1",
				BranchName:     "Branch One",
				BranchStreet:   "LTK",
				BranchWard:     "Ward 11",
				BranchDistrict: "District 10",
				BranchProvince: "HCMC",
				CreatedAt:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.Local),
				Manager:        "staff-1",
				OpenTime:       "07:00",
				CloseTime:      "21:00",
			},
		},
		{
			name: "Should return error when branch id not found",
			args: args{
				branchId: "branch-3",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			got, err := s.GetBranchDetail(ctx, tt.args.branchId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranchDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranchDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_branchServiceController_GetBranchStaff(t *testing.T) {
	type args struct {
		branchId string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Should get branch staffs correctly",
			args: args{
				branchId: "branch-1",
			},
			want: []string{"staff-1", "staff-2", "staff-3"},
		},
		{
			name: "Should return error when branch not exists",
			args: args{
				branchId: "branch-not-found",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branchServiceController{
				repository: testRepository,
			}
			got, err := s.GetBranchStaff(ctx, tt.args.branchId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranchStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranchStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
