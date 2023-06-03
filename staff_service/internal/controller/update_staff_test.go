package controller

import (
	"context"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_UpdateStaff(t *testing.T) {
	type args struct {
		request *schema.UpdateStaffRequest
		staffId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update staff successfully",
			args: args{
				request: &schema.UpdateStaffRequest{
					Name:     "HTTN",
					Street:   "Ly Thuong Kiet",
					Province: "Binh Duong",
				},
				staffId: "staff-1",
			},
		},
		{
			name: "Should return error when db update staff fails",
			args: args{
				request: &schema.UpdateStaffRequest{
					Name:     "HTTN",
					Street:   "Ly Thuong Kiet",
					Province: "Binh Duong",
				},
				staffId: "invalid-staff",
			},
			wantErr: true,
		},
		{
			name: "Should return error when request role invalid",
			args: args{
				request: &schema.UpdateStaffRequest{
					Name:     "HTTN",
					Street:   "Ly Thuong Kiet",
					Province: "Binh Duong",
					Role:     "invalid",
				},
				staffId: "staff-1",
			},
			wantErr: true,
		},
		{
			name: "Should return error when account adapter fails",
			args: args{
				request: &schema.UpdateStaffRequest{
					Name:     "HTTN",
					Street:   "Ly Thuong Kiet",
					Province: "Binh Duong",
					Role:     "3",
				},
				staffId: "invalid-account",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				accountAdapter: testAccount,
				repository:     testRepository,
			}
			if err := s.UpdateStaff(ctx, tt.args.request, tt.args.staffId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
