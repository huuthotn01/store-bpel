package controller

import (
	"context"
	"testing"
)

func Test_staffServiceController_DeleteStaff(t *testing.T) {
	type args struct {
		staffId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should delete staff successfully",
			args: args{
				staffId: "staff-1",
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				repository: testRepository,
			}
			if err := s.DeleteStaff(ctx, tt.args.staffId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
