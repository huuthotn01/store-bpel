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
